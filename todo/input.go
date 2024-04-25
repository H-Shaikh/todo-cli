package todo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ReadATodo() Todo {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n Enter the title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("\n Enter the description: ")
	desc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return Todo{
		Title:     strings.Trim(title, "\n"),
		Desc:      strings.Trim(desc, "\n"),
		Done:      false,
		CreatedAt: time.Now(),
	}

}

func InputLoop(existingTodos []Todo) []Todo {

	todos := make([]Todo, 0)
	for {
		fmt.Printf("\nPlease choose an option:\n")
		fmt.Printf("\n 1. Create a new todo \n")
		fmt.Printf("\n 2. Select a new todo \n")

		fmt.Printf("\n 0. Exit the loop\n")
		fmt.Printf("Enter your choice: ")

		var input string
		_, err := fmt.Scanf("%s", &input)

		if err != nil {
			fmt.Printf("E#1VUOMM - %v\n", err)
		}

		if input == "0" {
			break
		} else if input == "1" {
			t := ReadATodo()
			t.Print()
			todos = append(todos, t)
		} else if input == "2" {
			for i, t := range existingTodos {
				t.PrintIdAndTitle(i + 1)
			}
			var selectedTodoNumber int
			fmt.Printf("\nPlease enter the id of todo to select: ")
			_, err := fmt.Scanf("%v", &selectedTodoNumber)
			if err != nil {
				fmt.Printf("E#1VUQF1 - %v\n", err)
			}

			selectedTodoIndex := selectedTodoNumber - 1

			options := SingleTodoActions(existingTodos[selectedTodoIndex])
			switch options {
			case "1":
				fmt.Printf("You want to mark the todo %v as done", existingTodos[selectedTodoIndex].Title)
			case "2":
				fmt.Printf("You want to edit the title of Todo %v", existingTodos[selectedTodoIndex].Title)
			default:
				fmt.Printf("Invalid Selection!")
			}
		} else {
			fmt.Printf("You selected an invalid operaion try again")
		}
	}
	return todos
}

func SingleTodoActions(todo Todo) string {
	fmt.Printf("\n You have selected: %v", todo.Title)
	fmt.Printf("\n1. Mark as done")
	fmt.Printf("\n2. Edit the title")
	fmt.Printf("\n3. Please select the option: ")
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		log.Fatalln("Fatal error: ", err)

	}

	return input
}
