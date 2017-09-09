package store

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

var (
	NotFound       = errors.New("Not found")
	NotImplemented = errors.New("Not implemented")
)

type Store interface {
	Migrate() error
	NoteStore
}

type NoteStore interface {
	PutNote(note string) error
	GetNote(id string) (Note, error)
	GetNotes() ([]Note, error)
}

type impl struct {
	db *sql.DB
}

func (s impl) Migrate() error {
	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS notes ( id VARCHAR PRIMARY KEY, data VARCHAR, created TIMESTAMP)")
	return err

}

func (s impl) PutNote(note string) error {
	return NotImplemented
}

func (s impl) GetNote(id string) (Note, error) {
	//age := 21
	//rows, err := s.db.Query("SELECT name FROM users WHERE age = $1", age)
	return nil, NotFound
}

func (s impl) GetNotes() ([]Note, error) {
	return []Note{}, nil
}

func New(url string) (Store, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return impl{db}, err

}
