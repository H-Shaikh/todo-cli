package main

import (
	"encoding/json"
	"fmt"
	"github.com/H-Shaikh/todo-cli/todo"
	"os"
)

func main() {
	fmt.Println("This is a CLI app to manage todo lists. it is written in golang")

	f, err := os.OpenFile("todos.json",
		os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Printf("Error when opening the file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Some fatal error: %v", err)

		}
	}(f)
	// Read the contents of the existing file .......
	dat, err := os.ReadFile("todos.json")
	//fmt.Println(string(dat))

	var existingTodos []todo.Todo

	// ..... and try to cast it in a slice of todos
	err = json.Unmarshal(dat, &existingTodos)
	if err != nil {
		fmt.Printf("Fatal error when reading file: %v", err)
	}

	// ...... and print them!
	for _, t := range existingTodos {
		t.Print()
		fmt.Printf("######################\n")
	}

	//return

	// to read new todos and save them in  todos.json file
	todos := todo.InputLoop(existingTodos)
	for _, t := range todos {
		t.Print()
		fmt.Printf("######################\n")
	}

	// let's convert list of todos to JSON

	todoJsonBytes, err := json.MarshalIndent(todos, "", "	")
	if err != nil {
		fmt.Printf("\nSomething went wrong while encoding or marshalling to JSON: %v\n", err)
		return
	}
	todoJsonContents := string(todoJsonBytes)

	_, err = f.WriteString(todoJsonContents)
	if err != nil {
		fmt.Printf("\nError when writing to file: %v", err)
	}
	fmt.Printf("\nWrote content to file")

}
