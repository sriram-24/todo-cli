package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"todo-cli/models"
)

func Update(id string, todoUpdated string) error  {
	workspace, err:= GetPath()
	if err != nil {
		return err
	}
	todoPath:= filepath.Join(workspace,models.TODO_FILENAME)
	if _,fileerr := os.Stat(todoPath); os.IsNotExist(fileerr){
		return fmt.Errorf("failed to read filepath")
	} 	
	jsonData, openError:= os.ReadFile(todoPath)
	if openError !=nil {
		return openError
	}
	
	if len(jsonData) ==0 {
		return fmt.Errorf("todo list is empty")
	}

	var data []models.Todo
	jsonError := json.Unmarshal(jsonData,&data)
	
	if jsonError != nil {
		return fmt.Errorf("failed to unmarshal json: %s", jsonError)
	}
	
	for index, todo := range data {
		if todo.Id.String() == id {
			data[index].Todo = todoUpdated
			data[index].ModifiedAt = time.Now().Format("2006-01-02 15:04:05") 
			break
		}
	}
		
	jsonDataToWrite, jsonErrWrite := json.MarshalIndent(data, "", "  ")
	if jsonErrWrite != nil {
		return  jsonErrWrite
	}

	// Write json to file
	fileerr := os.WriteFile(todoPath, jsonDataToWrite, 0600)

	if fileerr != nil {
		return fmt.Errorf("error creating config file: %s", fileerr)
	}
	fmt.Printf("%s updated\n",id)
	return nil


}
