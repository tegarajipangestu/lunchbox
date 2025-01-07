package main

import (
	"fmt"
	"os"

	"github.com/tegarajipangestu/lunchbox/lunchbox"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [server|client]")
		return
	}

	if os.Args[1] == "server" {
		lunchbox.StartServer()
	} else if os.Args[1] == "client" {
		lunchbox.StartClient()
	} else {
		fmt.Println("Invalid argument. Use 'server' or 'client'")
	}
}
