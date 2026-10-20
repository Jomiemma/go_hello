package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
)

func main() {

	namePtr := flag.String("name", "World", "name to greet")

	flag.Parse()

	sessionID := uuid.New()

	fmt.Printf("Hello, %s!\n", *namePtr)
	fmt.Printf("Session UUID: %s\n", sessionID)

}
