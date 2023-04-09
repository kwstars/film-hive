package event

import "context"

// Event is the interface for all events.
type Event interface {
	Key() string
	Value() []byte
	//Header() map[string]string
	//Ack() error
	//Nack() error
}

// Handler is the interface for all event handlers.
type Handler func(context.Context, Event) error

// Sender is the interface for all event senders.
type Sender interface {
	Send(ctx context.Context, msg Event) error
	Close() error
}

// Receiver is the interface for all event receivers.
type Receiver interface {
	Receive(ctx context.Context, handler Handler) error
	Close() error
}
