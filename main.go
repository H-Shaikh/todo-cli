package main

import (
	"encoding/json"
	"fmt"
	"github.com/H-Shaikh/todo-cli/todo"
	"time"
)

func main() {
	fmt.Println("This is a CLI app to manage todo lists. it is written in golang")

	t := todo.Todo{
		Title:     "something",
		Desc:      "lorem ipsum lorem ipsum lorem ipsum lorem ipsum",
		Done:      true,
		CreatedAt: time.Now(),
	}
	if t.Validate() != nil {
		fmt.Printf("\nTodo is not valid. Error: %v\n", t.Validate())
	} else {
		fmt.Printf("\nTodo is valid \n")
	}

	todoBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("\nSomething went wrong while encoding to JSON: %v\n", err)

	}
	fmt.Printf("\nJSON Bytes: %v\n", todoBytes)
	fmt.Printf("\nJSON String: %v\n", string(todoBytes))

	t.Print()
}
