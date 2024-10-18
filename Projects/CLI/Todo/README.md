# Todo CLI Application

This is a Command-Line Interface (CLI) application for managing your to-dos, built with Golang. It allows you to add, remove, edit, list, and complete tasks directly from your terminal.

## Features

- Add new to-dos
- List all to-dos in a tabular format
- Mark to-dos as completed
- Delete to-dos
- Edit to-dos
- Track creation and completion times

## Packages Used

- `flag`: To handle command-line arguments.
- `aquasecurity/table`: To display to-dos in a tabular format.
- `os`: For file operations.
- `encoding/json`: For JSON utilities, like converting to-dos to bytes for storage.
- `time`: To manage creation and completion times of to-dos.
- `strconv`: To convert index values to strings for display in the table.

## Run the .exe/binary file and follow the instructions

```bash
    ./todo-cli -h for help
    ./todo-cli -list to list all todos
    ./todo-cli -add 'todo item' to add a new todo
    ./todo-cli -toggle 'index' to toggle a todo
    ./todo-cli -edit 'index' 'new todo item' to edit a todo
    ./todo-cli -delete 'index' to delete a todo
```

N.B: if you directly run the binary without specifying any action it will show Invalid Command as it can't match any actions.
