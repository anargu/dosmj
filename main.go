package main

import (
	"os"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	if port == "" {
		panic("env variables should be provided")
	}

	runServer()
}
