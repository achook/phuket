package main

import (
	"fmt"
	"phuket/logging"
	"phuket/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		logging.LogFatal(err)
	}

	fmt.Println(srv)

	fmt.Println("Hello, World!")
}
