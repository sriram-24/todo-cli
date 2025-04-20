/* Cli playground todo app
* This is a simple todo list example program written in go.
* Features:
*	- help - returns help in terminal
 */
package main

import (
	"fmt"
	"os"
	"todo-cli/commands"
)

// declare constants
var HELP_STRING string = `Welcome to Todo Cli!
 Todo Cli is a tool for creating todo using cli interface.
 
 Usage:
	 todo <command> [argument]
 
 Commands: 
	 add		adds a todo item
 `
var HELP_ADD string = `Expecting an argument: run 'todo-cli help' for help`

func main() {
	arg1 := os.Args[1]
	switch arg1 {
	case "help":
		fmt.Println(HELP_STRING)
	case "add":
		// reading argument 2

		if len(os.Args) == 3 {
			todo := os.Args[2]
			_, err := commands.Add(todo)
			if err != nil {
				fmt.Printf("Error occured: %s", err)
			}
		} else {
			fmt.Println(HELP_ADD)
		}
	}
}
