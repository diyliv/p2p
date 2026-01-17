package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <listen-port> ./app 8080 (for example)\n", os.Args[0])
		os.Exit(1)
	}

	var port int
	_, err := fmt.Sscanf(os.Args[1], "%d", &port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}
}
