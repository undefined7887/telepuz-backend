package network

import (
	"context"
	"fmt"
)

type Conn interface {
	SetData(value interface{})
	Data() interface{}

	Handle(path string, handler EventHandler)
	Send(path string, event Event)
}

type Event interface {
	fmt.Stringer
}

type EventHandler interface {
	NewEvent() Event
	Handle(ctx context.Context, event Event)
}
