package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo-cli/models"
)

func SetPath(path string) (nil, err error) {
	// check the path paramenter is valid or not.
	dirPath, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("directory does not exist: %s", path)
		} else {
			return nil, fmt.Errorf("error accessing directory: %s", path)

		}
	}
	// check if the path is a directory or not
	if dirPath.IsDir() {
		userconfigDir, err := os.UserConfigDir()

		if err != nil {
			return nil, err
		}

		// Create config application directory.
		appConfigPath := filepath.Join(userconfigDir, models.CONFIG_APPDIR)
		if err := os.MkdirAll(appConfigPath, 0755); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %w", err)
		}

		configFilePath := filepath.Join(appConfigPath, models.CONFIG_FILENAME)

		// Create app config and store it as json file
		appConfig := models.AppConfig{
			Workspace: path,
		}
		jsonData, err := json.MarshalIndent(appConfig, "", "  ")
		if err != nil {
			return nil, err
		}
		// write the app config json to file.
		fileerr := os.WriteFile(configFilePath, jsonData, 0600)
		if fileerr != nil {
			return nil, fmt.Errorf("error creating config file: %s", fileerr)
		}

		return nil, nil
	} else {
		return nil, fmt.Errorf("path exists but not a directory")
	}

}
