package main

import (
	"fmt"
	"log"

	"encoding/json"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	output, jsonerr := json.Marshal(messages)

	if jsonerr != nil {
		fmt.Printf("Error: %s", jsonerr.Error())
	} else {
		fmt.Println(string(output))
	}
}
