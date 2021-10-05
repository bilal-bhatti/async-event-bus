package events

import (
	"fmt"

	"github.com/bilal-bhatti/async-event-bus/util"
	"github.com/gofrs/uuid"
)

type Something struct {
	ID             string
	Parents        []string
	TimeToComplete int
}

func NewSomething() Something {
	id := uuid.Must(uuid.NewV4())
	return Something{
		ID:             id.String(),
		Parents:        make([]string, 0),
		TimeToComplete: util.RandomInRange(1, 20),
	}
}

func NewSomethingWithParent(parent Something) Something {
	st := NewSomething()
	st.Parents = append(parent.Parents, parent.ID)
	return st
}

func (s Something) ToString() string {
	return fmt.Sprintf("id: %s, chain: %v, ttc: %d", s.ID, s.Parents, s.TimeToComplete)
}
