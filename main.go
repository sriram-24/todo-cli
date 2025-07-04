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
	"todo-cli/cmd"
)

// declare constants
const HELP_STRING string = `Welcome to Todo Cli!
 todo-cli is a tool for creating todo using cli interface.
 
 Usage:
	 todo-cli <command> [argument]
 
 Commands: 
	 add		Add a todo item. eg: todo-cli add "Welcome to todo cli"
	 update 	Update an existing todo item eg: todo-cli update  "6d5d3876-8468-4416-82e4-c38ba097edd4" "new updated todo" 
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
			_, err := cmd.Add(todo)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(HELP_ADD)
		}
	case "list":
		// Returns list of todos
		todoList, listErr := cmd.ListAll()
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
			_, err := cmd.SetPath(path)
			if err != nil {
				fmt.Printf("Error occured: %s\n", err)
			}
		} else {
			fmt.Println(HELP_PATH)
		}
	case "get-path":
		fmt.Println(cmd.GetPath())
	case "delete":
		if len(os.Args) == 3{
			todo:= os.Args[2]
			err:= cmd.Delete(todo)
			if err != nil {
				fmt.Printf("error occured : %s",err)
			}
		}
	case "update":
		if len(os.Args) ==2 {
			fmt.Println("Expecting an argument of Id and Todo to be passed. Run todo-cli help for more information")
			return
		}
		if len(os.Args) ==3 {
			fmt.Println("Expecting an argument of Todo to be passed. Run todo-cli help for more information")
			return
		}
		if len(os.Args) == 4{
			todoId:= os.Args[2]
			todo:= os.Args[3]
			err:= cmd.Update(todoId,todo)
			if err != nil {
				fmt.Printf("error occured : %s",err)
			}
		}
	default:
		fmt.Println("Invalid command run todo-cli help for more information")
	}

}
