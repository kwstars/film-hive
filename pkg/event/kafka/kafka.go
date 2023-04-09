package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/kwstars/film-hive/pkg/event"
)

var (
	_ event.Sender   = (*Sender)(nil)
	_ event.Receiver = (*Receiver)(nil)
	_ event.Event    = (*Message)(nil)
)

type Sender struct {
	producer *kafka.Producer
	topic    string
}

func NewSender(broker, topic string) (event.Sender, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, err
	}
	s := &Sender{
		producer: p,
		topic:    topic,
	}

	return s, nil
}

func (s *Sender) Send(ctx context.Context, msg event.Event) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)
	if err := s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.topic, Partition: kafka.PartitionAny},
		Key:            []byte(msg.Key()),
		Value:          msg.Value(),
	}, deliveryChan); err != nil {
		return err
	}

	select {
	case ev := <-deliveryChan:
		m := ev.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return m.TopicPartition.Error
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Sender) Close() error {
	s.producer.Flush(15 * 1000)
	s.producer.Close()
	return nil
}

type Receiver struct {
	consumer *kafka.Consumer
	topic    string
}

func NewReceiver(broker, groupID, topic string) (event.Receiver, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
		//"security.protocol" : "PLAINTEXT",
		//"sasl.mechanism" : "PLAIN",
		//"sasl.username" : "XXX",
		//"sasl.password" : "XXX",
	})
	if err != nil {
		return nil, err
	}
	if err = c.SubscribeTopics([]string{topic}, nil); err != nil {
		return nil, err
	}
	return &Receiver{consumer: c, topic: topic}, nil
}

func (r *Receiver) Receive(ctx context.Context, handler event.Handler) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			ev := r.consumer.Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				if err := handler(ctx, &Message{msg: e}); err != nil {
					return err
				}
			case kafka.Error:
				return kafka.NewError(e.Code(), e.String(), e.IsFatal())
			}
		}
	}
}

func (r *Receiver) Close() error {
	return r.consumer.Close()
}

type Message struct {
	msg *kafka.Message
}

func (e *Message) Key() string {
	return string(e.msg.Key)
}

func (e *Message) Value() []byte {
	return e.msg.Value
}

func NewEvent(key string, value []byte) *Message {
	return &Message{
		msg: &kafka.Message{
			Key:   []byte(key),
			Value: value,
		},
	}
}
