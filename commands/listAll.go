package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo-cli/models"
)

func ListAll() ([]models.Todo, error) {
	workspaceDir, err := GetPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get path: %s", err)
	}
	filePath := filepath.Join(workspaceDir, models.TODO_FILENAME)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, nil
	}

	jsonData, fileErr := os.ReadFile(filePath)
	if fileErr != nil {
		return nil, fmt.Errorf("failed to read config file: %s", fileErr)
	}
	var data []models.Todo
	if len(jsonData) == 0 {
		return data, nil
	}
	jsonErr := json.Unmarshal(jsonData, &data)

	if jsonErr != nil {
		return nil, fmt.Errorf("failed to Unmarshal json: %s", jsonErr)
	}

	return data, nil
}
