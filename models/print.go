package models

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type Print struct {
	ID      string
	Message string
}

func NewPrint(message string) Print {
	id := uuid.Must(uuid.NewV4())
	return Print{
		ID:      id.String(),
		Message: message,
	}
}

func (s Print) ToString() string {
	return fmt.Sprintf("id: %s, message: %s", s.ID, s.Message)
}
