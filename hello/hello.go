package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Slice of names
	names := []string{"Tomas", "Daniela", "Angel", "Valentino"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	var result = ""

	for _, msg := range messages {
		result = result + msg + "\n"
	}

	fmt.Println(result)
}
