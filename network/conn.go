package network

import (
	"context"
	"fmt"
)

type Conn interface {
	Handle(path string, handler EventHandler)
	Send(path string, event Event)
}

type Event interface {
	fmt.Stringer
}

type EventHandler interface {
	NewEvent() Event
	ServeEvent(ctx context.Context, event Event)
}
