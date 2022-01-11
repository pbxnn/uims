package kafka

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type Message struct {
	Data        json.RawMessage   `json:"data"`
	SpanContext trace.SpanContext `json:"spanContext"`
}

func Send(ctx context.Context, kp sarama.SyncProducer, msg *sarama.ProducerMessage) (int32, int64, error) {
	d, err := msg.Value.Encode()
	if err != nil {
		return 0, 0, err
	}

	m := Message{
		Data:        d,
		SpanContext: trace.SpanContextFromContext(ctx),
	}
	v, err := json.Marshal(m)
	if err != nil {
		return 0, 0, err
	}

	msg.Value = sarama.ByteEncoder(v)

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
