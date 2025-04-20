package commands

import "fmt"

func Add(todo string) (nil, err error) {
	if todo == "" {
		return nil, fmt.Errorf("pass the todo string to add")
	}
	fmt.Printf("%s added", todo)
	return nil, nil
}
