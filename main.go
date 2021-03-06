package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usage(os.Args[0])
	}
	switch os.Args[1] {
	case "migrate":
		migrate()
	case "server":
		server_()
	case "get":
		get(os.Args[2:])
	case "list":
		list()
	case "put":
		put(strings.Join(os.Args[2:], " "))
	default:
		usage(os.Args[0])
	}
}
