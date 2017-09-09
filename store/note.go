package store

import (
	"time"

	"github.com/satori/go.uuid"
)

type Note interface {
	ID() string
	Data() string
	Created() int64
}

type noteImpl struct {
	id      uuid.UUID
	data    string
	created int64
}

func (i noteImpl) ID() string {
	return i.id.String()
}

func (i noteImpl) Data() string {
	return i.data
}

func (i noteImpl) Created() int64 {
	return i.created
}

func NewNote(data string) Note {
	return noteImpl{uuid.NewV4(), data, time.Now().UnixNano()}
}
