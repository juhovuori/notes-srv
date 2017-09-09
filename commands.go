package main

import (
	"fmt"
	"os"

	"github.com/juhovuori/minitwitter-srv/store"
)

func usage(cmd string) {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s migrate\n", cmd)
	fmt.Printf("  %s http\n", cmd)
	fmt.Printf("  %s list\n", cmd)
	fmt.Printf("  %s get [<id> ...]\n", cmd)
	fmt.Printf("  %s put <message>\n", cmd)
	os.Exit(1)
}

func mustGetStore() store.Store {
	url := "postgres://postgres:mintwitter@localhost:5432/?sslmode=disable"
	store, err := store.New(url)
	if err != nil {
		fmt.Printf("Could not open store: %s\n", err)
		os.Exit(1)
	}
	return store
}

func migrate() {
	store := mustGetStore()
	err := store.Migrate()
	if err != nil {
		fmt.Printf("Could not run migrations: %s\n", err)
		os.Exit(1)
	}
}

func http() {
	fmt.Printf("Not implemented\n")
	os.Exit(1)
}

func list() {
	store := mustGetStore()
	notes, err := store.GetNotes()
	if err != nil {
		fmt.Printf("Failed to get notes: %s\n", err)
		os.Exit(1)
	}
	for _, note := range notes {
		fmt.Println(note)
	}
}

func get(ids []string) {
	store := mustGetStore()
	status := 0
	for _, id := range ids {
		note, err := store.GetNote(id)
		if err != nil {
			fmt.Printf("Failed to get %s: %s\n", id, err)
			status = 1
		} else {
			fmt.Println(note)
		}
	}
	os.Exit(status)
}

func put(data string) {
	store := mustGetStore()
	note, err := store.PutNote(data)
	if err != nil {
		fmt.Printf("Failed to put: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Put: %v\n", note)
}
