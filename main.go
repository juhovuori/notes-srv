package main

import (
	"fmt"
	"os"

	"github.com/juhovuori/minitwitter-srv/store"
)

func main() {
	url := "postgres://postgres:mintwitter@localhost:5432/?sslmode=disable"
	store, err := store.New(url)
	if err != nil {
		fmt.Printf("Could not open store: %s\n", err)
		os.Exit(1)
	}
	err = store.Migrate()
	if err != nil {
		fmt.Printf("Could not run migrations: %s\n", err)
		os.Exit(1)
	}
	return
}
