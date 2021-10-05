package handlers

import (
	"context"
	"fmt"

	"github.com/bilal-bhatti/async-event-bus/concur"
	"github.com/bilal-bhatti/async-event-bus/events"
	"github.com/mustafaturan/bus/v3"
)

func DoPrint(ctx context.Context, emit concur.Emitter, e bus.Event) {
	t := e.Data.(events.Print)
	fmt.Printf("print: %s\n", t.ToString())
}
