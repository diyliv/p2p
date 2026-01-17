package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s -port <listen-port> [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("\nExample:")
		fmt.Printf("  %s -port 8080\n", os.Args[0])
		fmt.Printf("  %s -port 8080 -algorithm RSA -key-size 4096\n", os.Args[0])
	}
}

func main() {
	// move to config file
	var (
		port      int
		algorithm string
		keySize   int
	)

	flag.IntVar(&port, "port", 0, "Provide port on which service is going to run (must have)")
	flag.StringVar(&algorithm, "algorithm", "RSA", "Encryption algorithm (supported RSA)")
	flag.IntVar(&keySize, "key-size", 2048, "Key size for asymmetric algorithms")
	flag.Parse()

	if port == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if port < 1 || port > 65535 {
		log.Fatalf("Invalid port number %d: port must be between 1 and 65535", port)
	}

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <listen-port> ./app 8080 (for example)\n", os.Args[0])
		os.Exit(1)
	}

	_, err := fmt.Sscanf(os.Args[1], "%d", &port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}
}
