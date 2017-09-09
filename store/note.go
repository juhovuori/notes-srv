package store

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type Note interface {
	ID() string
	Data() string
	Created() time.Time
}

type noteImpl struct {
	id      uuid.UUID
	data    string
	created time.Time
}

func (i noteImpl) ID() string {
	return i.id.String()
}

func (i noteImpl) Data() string {
	return i.data
}

func (i noteImpl) Created() time.Time {
	return i.created
}

func (i noteImpl) String() string {
	return fmt.Sprintf("%s(%v): %s", i.ID(), i.Created(), i.Data())
}

func NewNote(data string) Note {
	return noteImpl{uuid.NewV4(), data, time.Now()}
}
