## 📝 ToDoGoru - CLI Task Manager in Go ##

ToDoGoru is a simple command-line interface (CLI) task manager built in **Go (Golang)**.  
It allows users to manage their tasks directly from the terminal by **adding**, **listing**, **marking as complete**, and **deleting** tasks.  
This app is designed to be **easy to use** and **efficient** for managing everyday tasks.

---

## ✨ Features

- ✅ **Add tasks**: Create tasks by typing `todoguru add "task"`.
- ✅ **Mark tasks as complete**: Mark a task as done with `todoguru done <taskID>`.
- ✅ **List tasks**: View all tasks with `todoguru list`.
- ✅ **Delete tasks**: Remove tasks with `todoguru delete <taskID>`.

---

## ⚙️ Prerequisites

- Go **1.16** or higher installed on your system.
- Basic knowledge of Go and CLI applications.

---

## 🚀 Installation

### 1. Clone the repository

```bash
git clone https://github.com/GuruOrGoru/ToDoGoru.git
```

### 2. Navigate to the project directory

```bash
cd ToDoGoru
```

### 3. Build the project

```bash
go build -o todoguru
```

This will create a `todoguru` binary in the current directory.

### 4. (Optional) Install globally

```bash
sudo mv todoguru /usr/local/bin/
```

---

## 📖 Usage

### 1. Add a Task

```bash
todoguru add "Task description here"
```

**Example:**
```bash
todoguru add "Buy groceries"
```

### 2. List All Tasks

```bash
todoguru list
```

This will display a list of all tasks with their IDs, descriptions, and completion statuses.

### 3. Mark a Task as Complete

```bash
todoguru done <taskID>
```

**Example:**
```bash
todoguru done 1
```

### 4. Delete a Task

```bash
todoguru delete <taskID>
```

**Example:**
```bash
todoguru delete 2
```

---

## 🎯 Example Workflow

```bash
# Add tasks
todoguru add "Buy groceries"
todoguru add "Complete Go project"

# List all tasks
todoguru list
# Output:
# 1. Buy groceries - Incomplete
# 2. Complete Go project - Incomplete

# Mark a task as done
todoguru done 1

# List all tasks again
todoguru list
# Output:
# 1. Buy groceries - Complete
# 2. Complete Go project - Incomplete

# Delete a task
todoguru delete 2

# List all tasks again
todoguru list
# Output:
# 1. Buy groceries - Complete
```

---

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-name`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add new feature'`)
5. Push to the branch (`git push origin feature-name`)
6. Open a pull request

---

## developers: 
Siddhartha Dhakal