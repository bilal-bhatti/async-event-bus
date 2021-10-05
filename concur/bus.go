package concur

import (
	"context"
	"sync"

	"github.com/mustafaturan/bus/v3"
	"github.com/mustafaturan/monoton/v2"
	"github.com/mustafaturan/monoton/v2/sequencer"
)

type AsyncBus struct {
	bus *bus.Bus
	wg  *sync.WaitGroup
}

func NewBus() (*AsyncBus, error) {
	node := uint64(1)
	initialTime := uint64(1577865600000)
	monoton, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		return nil, err
	}

	var idGenerator bus.Next = monoton.Next

	b, err := bus.NewBus(idGenerator)
	if err != nil {
		return nil, err
	}

	return &AsyncBus{
		bus: b,
		wg:  &sync.WaitGroup{},
	}, nil
}

func (b AsyncBus) Shutdown(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		b.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (b AsyncBus) RegisterTopics(topics ...string) {
	b.bus.RegisterTopics(topics...)
}

type Handler func(context.Context, Emitter, bus.Event)

func (b AsyncBus) RegisterHandler(name string, handler Handler, matcher string) {
	b.bus.RegisterHandler(name, bus.Handler{Handle: b.wrap(handler), Matcher: matcher})
}

func (b AsyncBus) wrap(handler Handler) func(ctx context.Context, event bus.Event) {
	return func(ctx context.Context, event bus.Event) {
		go func() {
			defer b.wg.Done()
			handler(ctx, b, event)
		}()
	}
}

type Emitter interface {
	Emit(ctx context.Context, topic string, data interface{}) error
}

func (b AsyncBus) Emit(ctx context.Context, topic string, data interface{}) error {
	b.wg.Add(1)
	err := b.bus.Emit(ctx, topic, data)
	if err != nil {
		return err
	}

	return nil
}
