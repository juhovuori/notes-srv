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

func get(ids []string) {
	fmt.Printf("Not implemented\n")
	os.Exit(1)
}

func put(data string) {
	fmt.Printf("Put %s\n", data)
	store := mustGetStore()
	note, err := store.PutNote(data)
	if err != nil {
		fmt.Printf("Failed to put: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Put: %v\n", note)
}
