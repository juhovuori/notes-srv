package main

import (
	"fmt"
	"os"

	"github.com/juhovuori/notes-srv/server"
	"github.com/juhovuori/notes-srv/store"
)

func usage(cmd string) {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s migrate\n", cmd)
	fmt.Printf("  %s server\n", cmd)
	fmt.Printf("  %s list\n", cmd)
	fmt.Printf("  %s get [<id> ...]\n", cmd)
	fmt.Printf("  %s put <message>\n", cmd)
	os.Exit(1)
}

func mustGetStore(cfg Config) store.Store {
	url := cfg.DB
	store, err := store.New(url)
	if err != nil {
		fmt.Printf("Could not open store: %s\n", err)
		os.Exit(1)
	}
	return store
}

func mustConfigure() Config {
	cfg, err := configure()
	if err != nil {
		fmt.Printf("Could not configure: %s\n", err)
		os.Exit(1)
	}
	return cfg
}

func migrate() {
	cfg := mustConfigure()
	store := mustGetStore(cfg)
	err := store.Migrate()
	if err != nil {
		fmt.Printf("Could not run migrations: %s\n", err)
		os.Exit(1)
	}
}

func server_() {
	cfg := mustConfigure()
	store := mustGetStore(cfg)
	server, err := server.New(store)
	if err != nil {
		fmt.Printf("Failed to create server: %s\n", err)
		os.Exit(1)
	}
	err = server.Start()
	if err != nil {
		fmt.Printf("Server failed: %s\n", err)
		os.Exit(1)
	}
}

func list() {
	cfg := mustConfigure()
	store := mustGetStore(cfg)
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
	cfg := mustConfigure()
	store := mustGetStore(cfg)
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
	cfg := mustConfigure()
	store := mustGetStore(cfg)
	note, err := store.PutNote(data)
	if err != nil {
		fmt.Printf("Failed to put: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Put: %v\n", note)
}
