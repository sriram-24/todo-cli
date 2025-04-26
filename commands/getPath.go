package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo-cli/models"
)

func GetPath() (string, error) {
	// Get user config dir
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	// Read the app config from file
	filePath := filepath.Join(userConfigPath, models.CONFIG_APPDIR, models.CONFIG_FILENAME)
	// check path exists or not
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("workspace config is missing. run 'todo-cli help' for more information")
	}

	jsonData, fileErr := os.ReadFile(filePath)
	if fileErr != nil {
		return "", fmt.Errorf("failed to read config file: %s", fileErr)
	}

	if len(jsonData) == 0 {
		return "", fmt.Errorf("failed to read config: set workspace path before adding todo")
	}

	// parse json and return workspace config as string.
	var data models.AppConfig
	jsonErr := json.Unmarshal(jsonData, &data)

	if jsonErr != nil {
		return "", fmt.Errorf("failed to Unmarshal json: %s", jsonErr)
	}

	path := data.Workspace
	return path, nil

}
