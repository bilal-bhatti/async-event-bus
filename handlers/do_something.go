package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/bilal-bhatti/async-event-bus/concur"
	"github.com/bilal-bhatti/async-event-bus/models"
	"github.com/mustafaturan/bus/v3"
)

func DoSomething(ctx context.Context, bus concur.AsyncBus, e bus.Event) {
	t := e.Data.(models.Something)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic event: %s, with error %v\n", t.ID, r)
		}
	}()

	if t.TimeToComplete == 17 {
		panic(fmt.Errorf("error 17"))
	}

	fmt.Printf("sleep: %s\n", t.ToString())
	// simulate time to do some work
	time.Sleep(time.Duration(t.TimeToComplete) * time.Second)
	fmt.Printf("awake: %s\n", t.ToString())

	if t.TimeToComplete == 5 || t.TimeToComplete == 10 || t.TimeToComplete == 15 {
		err := bus.Emit(ctx, "do.something", models.NewSomethingWithParent(t))
		if err != nil {
			panic(err)
		}
	}

	err := bus.Emit(ctx, "do.print", models.NewPrint(fmt.Sprintf("%v", t.ToString())))
	if err != nil {
		panic(err)
	}
}
