/* Cli playground todo app
* This is a simple todo list example program written in go.
* Features:
*	- help - returns help in terminal
 */
package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"todo-cli/commands"
)

// declare constants
const HELP_STRING string = `Welcome to Todo Cli!
 todo-cli is a tool for creating todo using cli interface.
 
 Usage:
	 todo-cli <command> [argument]
 
 Commands: 
	 add		Add a todo item. eg: todo-cli add "Welcome to todo cli"
	 remove		Removes a todo item by passing todo ID. eg: todo-cli remove "6d5d3876-8468-4416-82e4-c38ba097edd4"
	 list		List All todo items from workspace. eg: todo-cli list
	 set-path	Set workspace path to add notes. eg: todo-cli set-path "C:\notes"
	 get-path	Get current workspace path. eg: todo-cli get-path
 `
const HELP_ADD string = `Expecting an argument: run 'todo-cli help' for help`
const HELP_PATH string = `Expecting an argument: Path is not passed as argument`

func main() {
	// If no arguments return help
	if(len(os.Args) == 1){
		 fmt.Println(HELP_STRING)
		return
	}
	arg1 := os.Args[1]
	switch arg1 {
	case "help":
		fmt.Println(HELP_STRING)
	case "add":
		// reading argument 2

		if len(os.Args) == 3 {
			// reading third argument for Todo and add it.
			todo := os.Args[2]
			_, err := commands.Add(todo)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(HELP_ADD)
		}
	case "list":
		// Returns list of todos
		todoList, listErr := commands.ListAll()
		if listErr != nil {
			fmt.Println(listErr)
		}

		// Write in console as table format.
		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(writer, "Id\tTodo\tCreated At\tModified At")
		fmt.Fprintln(writer, strings.Repeat("-------------------\t", 4))

		for _, row := range todoList {
			fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", row.Id, row.Todo, row.CreatedAt, row.ModifiedAt)
		}

		err := writer.Flush()
		if err != nil {
			fmt.Println("Error flushing tabwriter:", err)
		}

	case "set-path":
		if len(os.Args) == 3 {
			path := os.Args[2]
			_, err := commands.SetPath(path)
			if err != nil {
				fmt.Printf("Error occured: %s", err)
			}
		} else {
			fmt.Println(HELP_PATH)
		}
	case "get-path":
		fmt.Println(commands.GetPath())
	case "delete":
		if len(os.Args) == 3{
			todo:= os.Args[2]
			err:= commands.Delete(todo)
			if err != nil {
				fmt.Printf("error occured : %s",err)
			}
		}
	}

}
