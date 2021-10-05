package handlers

import (
	"context"
	"fmt"

	"github.com/bilal-bhatti/async-event-bus/concur"
	"github.com/bilal-bhatti/async-event-bus/events"
	"github.com/bilal-bhatti/async-event-bus/util"
	"github.com/mustafaturan/bus/v3"
)

func DoPrint(ctx context.Context, emitter concur.Emitter, e bus.Event) {
	t := e.Data.(events.Print)
	fmt.Printf("print: %s\n", t.ToString())

	q := util.RandomInRange(0, 2)

	if q == 0 {
		err := emitter.Emit(ctx, "do.construct", ConstructEvent{Value: fmt.Sprintf("%v", t.ToString())})
		if err != nil {
			panic(err)
		}
	}
}
