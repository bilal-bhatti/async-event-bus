package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bilal-bhatti/async-event-bus/concur"
	"github.com/bilal-bhatti/async-event-bus/events"
	"github.com/bilal-bhatti/async-event-bus/handlers"
)

func main() {
	bus, err := concur.NewBus()
	if err != nil {
		panic(err)
	}

	bus.RegisterTopics("do.something", "do.print", "do.construct")

	bus.RegisterHandler("do.something.worker", handlers.DoSomething, "do.something")
	bus.RegisterHandler("do.print.worker", handlers.DoPrint, "do.print")
	constructor := handlers.NewConstruct("dbstring", "awsstring")
	bus.RegisterHandler("do.construct.worker", constructor.Handle(), "do.construct")

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go gracefulShutdown(bus, quit, done)

	ctx := context.Background()

	for i := 0; i < 15; i++ {
		e := events.NewSomething()
		fmt.Printf("emit:  %s\n", e.ToString())
		err := bus.Emit(ctx, "do.something", e)
		if err != nil {
			panic(err)
		}
	}

	<-done
	fmt.Println("shutdown complete")
}

func gracefulShutdown(bus *concur.AsyncBus, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	fmt.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := bus.Shutdown(ctx); err != nil {
		fmt.Printf("could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
