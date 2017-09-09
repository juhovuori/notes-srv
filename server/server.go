package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/juhovuori/notes-srv/store"
)

type Server interface {
	Start() error
}

type notes struct {
	Notes []note `json:"notes"`
}

type note struct {
	ID      string
	Created time.Time
	Data    string
}

type impl struct {
	store store.Store
}

func (s *impl) list() (notes, error) {
	res := notes{[]note{}}
	n, err := s.store.GetNotes()
	if err != nil {
		return res, err
	}
	for _, n := range n {
		res.Notes = append(res.Notes, note{n.ID(), n.Created(), n.Data()})
	}
	return res, nil
}

func (s *impl) get(id string) (note, error) {
	n, err := s.store.GetNote(id)
	if err != nil {
		return note{}, err
	}
	return note{n.ID(), n.Created(), n.Data()}, nil
}

func (s *impl) post(body io.Reader) (note, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return note{}, err
	}
	n, err := s.store.PutNote(string(data))
	if err != nil {
		return note{}, err
	}
	return note{n.ID(), n.Created(), n.Data()}, nil
}

func (s *impl) handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/"):]
	defer r.Body.Close()
	if r.Method == "GET" && id == "" {
		notes, err := s.list()
		sendResponse(w, notes, err)
		return
	}
	if r.Method == "GET" && id != "" {
		note, err := s.get(id)
		sendResponse(w, note, err)
		return
	}
	if r.Method == "POST" {
		note, err := s.post(r.Body)
		sendResponse(w, note, err)
		return
	}
}

func sendResponse(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(200)
	w.Write(res)
}

func (i *impl) Start() error {
	http.HandleFunc("/", i.handler)
	return http.ListenAndServe(":8080", nil)

}

func New(s store.Store) (Server, error) {
	return &impl{s}, nil
}
