package handlers

import (
	"context"
	"fmt"

	"github.com/bilal-bhatti/async-event-bus/concur"
	"github.com/bilal-bhatti/async-event-bus/models"
	"github.com/mustafaturan/bus/v3"
)

func DoPrint(ctx context.Context, bus concur.AsyncBus, e bus.Event) {
	t := e.Data.(models.Print)
	fmt.Printf("print: %s\n", t.ToString())
}
