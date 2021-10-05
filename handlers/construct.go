package handlers

import (
	"context"
	"fmt"

	"github.com/bilal-bhatti/async-event-bus/concur"
	"github.com/bilal-bhatti/async-event-bus/events"
	"github.com/mustafaturan/bus/v3"
)

// an example of using structs instead of pure funcs as event handlers
type Construct struct {
	DB  string
	AWS string
}

type ConstructEvent struct {
	Value string
}

func NewConstruct(db, aws string) *Construct {
	return &Construct{
		DB:  db,
		AWS: aws,
	}
}

func (c Construct) Handle() concur.Handler {
	return func(ctx context.Context, emitter concur.Emitter, e bus.Event) {
		t := e.Data.(ConstructEvent)
		// do something with DB and AWS
		// ...
		fmt.Printf("struc: %s\n", t.Value)
		// emit another event
		err := emitter.Emit(ctx, "do.print", events.NewPrint(fmt.Sprintf("%v", t.Value)))
		if err != nil {
			panic(err)
		}
	}
}
