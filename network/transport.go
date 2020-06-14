package network

import (
	"context"
	"fmt"
	"time"
)

type Transport interface {
	SetTimeout(timeout time.Duration)

	Handle(path string, handler Handler)
	HandleQueue(path, queue string, handler Handler)

	Send(path string, event Event)
}

type Handler interface {
	NewEvent() Event
	Handle(ctx context.Context, event Event)
}

type Event interface {
	fmt.Stringer

	Reply() Reply
}

type Reply interface {
	fmt.Stringer
}
