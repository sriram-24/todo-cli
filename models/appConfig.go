package models

const CONFIG_APPDIR string = "todo-cli"
const CONFIG_FILENAME string = "config.json"
const TODO_FILENAME string = "todo.json"

type AppConfig struct {
	Workspace string
}
