package gorm_plugin

import (
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	gormSpanKey        = "__gorm_span"
	callBackBeforeName = "tracing:before"
	callBackAfterName  = "tracing:after"
)

func newTracer() trace.Tracer {
	tracer := otel.Tracer("kratos")
	return tracer
}

func before(db *gorm.DB) {
	ctx := db.Statement.Context
	if _, ok := transport.FromServerContext(ctx); ok {
		tracer := newTracer()
		_, span := tracer.Start(ctx, db.Name(), trace.WithSpanKind(trace.SpanKindInternal))
		db.InstanceSet(gormSpanKey, span)
	}

	return
}

func after(db *gorm.DB) {
	// 从GORM的DB实例中取出span
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}

	// 断言进行类型转换
	span, ok := _span.(trace.Span)
	if !ok {
		return
	}

	ctx := db.Statement.Context
	if _, ok := transport.FromServerContext(ctx); !ok {
		return
	}

	defer span.End()

	// Error
	if db.Error != nil {
		span.RecordError(db.Error)
	}

	// sql
	span.SetAttributes(attribute.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	span.SetAttributes(attribute.Int64("rowsAffected", db.Statement.RowsAffected))

	return
}

type OpentracingPlugin struct{}

func (op *OpentracingPlugin) Name() string {
	return "tracingPlugin"
}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	db.Callback().Create().Before("gorm_plugin:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm_plugin:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm_plugin:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm_plugin:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm_plugin:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm_plugin:raw").Register(callBackBeforeName, before)

	// 结束后
	db.Callback().Create().After("gorm_plugin:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm_plugin:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm_plugin:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm_plugin:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm_plugin:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm_plugin:raw").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &OpentracingPlugin{}
