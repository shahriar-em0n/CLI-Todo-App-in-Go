# CLI Todo App in Go

A simple, elegant command-line todo application built with Go. Manage your tasks efficiently with a clean table-based interface and persistent JSON storage.

## ✨ Features

- ✅ Add new todos with titles
- 📝 Edit existing todos
- 🗑️ Delete todos by index
- 🔄 Toggle completion status
- 📋 List all todos in a formatted table
- 💾 Persistent storage using JSON
- 🎨 Beautifully formatted table output with emojis

## 🚀 Installation

### Prerequisites

- Go 1.18 or higher

### Clone and Build

```bash
# Clone the repository
git clone <your-repo-url>
cd CLI-Todo-App-in-Go

# Install dependencies
go mod download

# Build the application
go build -o todo-app
```

## 📖 Usage

The application supports the following commands:

### List All Todos

```bash
./todo-app -list
```

**Example Output:**
```
┌───┬───────────┬───────────┬───────────────────────────────┬───────────────────────────────┐
│ # │   Title   │ Completed │          Created At           │          Completed At         │
├───┼───────────┼───────────┼───────────────────────────────┼───────────────────────────────┤
│ 0 │ Buy milk  │ ✅        │ Sat, 06 Dec 2025 08:15:25 +06 │ Sat, 06 Dec 2025 08:38:45 +06 │
│ 1 │ Buy bread │ ❌        │ Sat, 06 Dec 2025 08:15:25 +06 │                               │
└───┴───────────┴───────────┴───────────────────────────────┴───────────────────────────────┘
```

### Add a New Todo

```bash
./todo-app -add "Your task here"
```

**Example:**
```bash
./todo-app -add "Complete Go project"
```

### Edit a Todo

Edit a todo by specifying its index and new title in the format `index:new_title`:

```bash
./todo-app -edit "0:Updated task title"
```

**Example:**
```bash
./todo-app -edit "0:Buy organic milk"
```

### Toggle Todo Completion

Toggle the completion status of a todo by index:

```bash
./todo-app -toggle 0
```

This will mark the todo as completed (✅) and record the completion time.

### Delete a Todo

Delete a todo by its index:

```bash
./todo-app -delete 0
```

## 📁 Project Structure

```
CLI-Todo-App-in-Go/
├── main.go         # Application entry point
├── command.go      # Command-line flag definitions and execution logic
├── todo.go         # Todo struct and methods (add, edit, delete, toggle, print)
├── storage.go      # Generic storage implementation for JSON persistence
├── todos.json      # Persistent storage file (created automatically)
├── go.mod          # Go module definition
├── go.sum          # Go module checksums
└── README.md       # This file
```

## 🛠️ Tech Stack

- **Language:** Go
- **Dependencies:**
  - [github.com/aquasecurity/table](https://github.com/aquasecurity/table) - Beautiful table formatting for CLI output

## 📝 Code Overview

### Core Components

#### `Todo` Struct
```go
type Todo struct {
    Title       string
    Completed   bool
    CreatedAt   time.Time
    CompletedAt *time.Time
}
```

#### `Storage` (Generic)
The storage module uses Go generics to provide type-safe JSON persistence:
- `Save(data T)` - Saves data to JSON file
- `Load(data *T)` - Loads data from JSON file

#### Command Flags
- `-add <title>` - Add a new todo
- `-list` - List all todos
- `-edit <id:title>` - Edit todo by index
- `-delete <id>` - Delete todo by index
- `-toggle <id>` - Toggle completion status

## 🎯 How It Works

1. **Storage Loading:** On startup, the app loads existing todos from `todos.json`
2. **Command Execution:** The app parses command-line flags and executes the requested operation
3. **Storage Saving:** After executing the command, changes are persisted back to `todos.json`

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is open source and available under the MIT License.

## 👨‍💻 Author

Mohammad Shahriar

---

**Happy Task Managing! 📝✨**