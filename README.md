# Todo CLI

A simple command-line todo list application built with Go and Cobra. This is a learning project for exploring Go development and CLI application patterns.

## Features

- ✅ Add new tasks
- 📝 Edit existing tasks  
- ✔️ Mark tasks as complete
- 🗑️ Delete tasks
- 📋 List tasks (with optional view of completed tasks)
- ⏰ Track task creation timestamps
- 💾 JSON-based storage

## Installation

### Prerequisites
- Go 1.25.0 or later

### Build from source
```bash
git clone https://github.com/spencerwcowles/golang-todo
cd todo
go build -o todo
```

### Install directly
```bash
go install
```

## Usage

### Basic Commands

#### Add a new task
```bash
todo add "Buy groceries"
todo new "Learn Go programming"  # alias
```

#### List tasks
```bash
todo list                    # Show incomplete tasks only
todo list --all             # Show all tasks including completed
todo list -a                # Short flag for --all
todo all                    # Alias for list
```

#### Complete a task
```bash
todo complete 1             # Mark task ID 1 as complete
todo done 2                 # Alias for complete
```

#### Edit a task
```bash
todo edit 1 "Updated task title"
todo e 1 "New description"   # Short alias
```

#### Delete a task
```bash
todo delete 1               # Delete task ID 1
todo del 2                  # Alias options:
todo remove 3               # del, remove, rm, rem
todo rm 4
```

### Example Workflow

```bash
# Add some tasks
todo add "Write README documentation"
todo add "Review Go best practices"
todo add "Deploy application"

# List current tasks
todo list
# Output:
# ID    Task                     Created
# 1     Write README documentation    just now
# 2     Review Go best practices      just now  
# 3     Deploy application           just now

# Complete the first task
todo complete 1

# Edit the second task
todo edit 2 "Study Go concurrency patterns"

# List all tasks including completed ones
todo list --all
# Output:
# ID    Task                          Created    Completed
# 1     Write README documentation    2m ago     true
# 2     Study Go concurrency patterns 2m ago     false
# 3     Deploy application           2m ago     false

# Delete a task
todo delete 3
```

## Data Storage

Tasks are stored in a `todo.json` file in the current directory. The JSON structure includes:

```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Sample task",
      "completed": false,
      "time": "2024-01-01T12:00:00Z"
    }
  ],
  "nextId": 2
}
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [timediff](https://github.com/mergestat/timediff) - Human-readable time differences

## Development

This project uses Go modules. To get started with development:

```bash
# Clone the repository
git clone https://github.com/spencerwcowles/golang-todo
cd todo

# Install dependencies
go mod download

# Run the application
go run main.go <command>

# Build the binary
go build -o todo
```

## Project Structure

```
.
├── main.go              # Application entry point
├── go.mod              # Go module definition
├── cmd/
│   ├── root.go         # Root command and shared types
│   ├── add.go          # Add task command
│   ├── list.go         # List tasks command
│   ├── complete.go     # Complete task command
│   ├── edit.go         # Edit task command
│   └── delete.go       # Delete task command
└── todo.json          # Task data storage (created on first run)
```

## Learning Notes

This project was built as a first Go application to explore:
- Go syntax and idioms
- CLI application development with Cobra
- JSON file handling
- Command-line argument parsing
- Project structure and organization

## Future Improvements

- [ ] Add task priorities
- [ ] Support for task categories/tags
- [ ] Due date functionality
- [ ] Data persistence options (SQLite, PostgreSQL)
- [ ] Import/export capabilities
- [ ] Task search functionality
