package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	fmt.Println("Starting Interactive NODUS Client + Server")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		serverMain() // Reference the server function you set up
	}()

	go func() {
		defer wg.Done()
		clientMain() // Reference the client function you set up
	}()

	wg.Wait()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Handle server logic, such as reading messages, here
	fmt.Fprintln(conn, "Hello from the server!")
}

func serverMain() {

	// Copy server logic from nodusd/main.go or refactor to a shared package
	fmt.Println("Starting NODUS Server...")
	ln, err := net.Listen("tcp", ":8080") // Choose an appropriate port
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn) // Define this function to handle each client
	}
}

func clientMain() {
	// Copy client logic from nodus-client/main.go or refactor to a shared package
	fmt.Println("Starting NODUS Client...")
	conn, err := net.Dial("tcp", "localhost:8080") // Match the server port
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Hello from the client!")
	// Read response from the server if needed
}
