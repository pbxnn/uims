package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

func Send(ctx context.Context, kp sarama.SyncProducer, msg *sarama.ProducerMessage) (int32, int64, error) {

	partition, offset, err := kp.SendMessage(msg)

	if _, ok := transport.FromServerContext(ctx); ok {
		tracer := otel.Tracer("kratos")
		_, span := tracer.Start(ctx, msg.Topic, trace.WithSpanKind(trace.SpanKindProducer), trace.WithTimestamp(time.Now()))
		defer span.End()

		msgKey, _ := msg.Key.Encode()
		msgValue, _ := msg.Value.Encode()
		attrs := []attribute.KeyValue{
			attribute.String("topic", msg.Topic),
			attribute.String("key", string(msgKey)),
			attribute.String("value", string(msgValue)),
			attribute.Int64("offset", offset),
			attribute.Int64("partition", int64(partition)),
		}

		if err != nil {
			attrs = append(attrs, attribute.String("error", err.Error()))
		}
		span.SetAttributes(attrs...)
	}

	return partition, offset, err
}
