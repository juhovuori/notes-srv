package store

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("Not found")
)

type Store interface {
	Migrate() error
	NoteStore
}

type NoteStore interface {
	PutNote(note string) (Note, error)
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

func (s impl) PutNote(data string) (Note, error) {
	note := NewNote(data)
	_, err := s.db.Exec("INSERT INTO notes (id, data, created) VALUES ($1, $2, $3)", note.ID(), note.Data(), note.Created())
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s impl) GetNote(ID string) (Note, error) {
	var note noteImpl
	row := s.db.QueryRow("SELECT id, data, created FROM notes WHERE id = $1", ID)
	err := row.Scan(&note.id, &note.data, &note.created)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return note, nil
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
