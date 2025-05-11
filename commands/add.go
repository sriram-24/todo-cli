package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo-cli/models"
)

func Add(todo string) (nil, err error) {
	if todo == "" {
		return nil, fmt.Errorf("pass the todo string to add")
	}
	// create new Todo
	newTodo := models.NewTodo(todo)
	workspaceDir, err := GetPath()
	if err != nil {
		return nil, err
	}

	// Get All todos and append new todo
	todoList, listErr := ListAll()
	if listErr != nil {
		return nil, listErr
	}
	todoList = append(todoList, newTodo)

	// Convert todo as json
	jsonData, jsonErr := json.MarshalIndent(todoList, "", "  ")
	if jsonErr != nil {
		return nil, jsonErr
	}

	// Write json to file
	todoPath := filepath.Join(workspaceDir, models.TODO_FILENAME)
	fileerr := os.WriteFile(todoPath, jsonData, 0600)

	if fileerr != nil {
		return nil, fmt.Errorf("error creating config file: %s", fileerr)
	}

	fmt.Printf("%s added\n", todo)
	return nil, nil
}
