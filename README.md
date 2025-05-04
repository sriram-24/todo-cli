# todo-cli
A simple command line todo application to add, list and delete the todos.

### Installing

Checkout the [releases](https://github.com/sriram-24/todo-cli/releases) and install based on your platform or you can build it from the source. Make sure you have `go` already installed on the system.

```bash
go build
```

### Usage

Run the help command to get the commands and how to use the application. 

```bash
todo-cli help
```

### Commands

- `add` - Add a todo item. eg: todo-cli add "Welcome to todo cli"
- `remove` - Removes a todo item by passing todo ID. eg: todo-cli remove "6d5d3876-8468-4416-82e4-c38ba097edd4"
- `list` - List All todo items from workspace. eg: todo-cli list
- `set-path` - Set workspace path to add notes. eg: todo-cli set-path "C:\notes"
- `get-path` - Get current workspace path. eg: todo-cli get-path

