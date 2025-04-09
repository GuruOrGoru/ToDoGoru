ToDoGoru - CLI Task Manager in Go

ToDoGoru is a simple command-line interface (CLI) task manager built in Go (Golang). It allows users to manage their tasks directly from the terminal by adding, listing, marking tasks as complete, and deleting them. This app is designed to be easy to use and efficient for managing everyday tasks.

🚀 Features

Add tasks: Create tasks by typing todoguru add "task".

Mark tasks as complete: Mark a task as done with todoguru done taskID.

List tasks: View all tasks with todoguru list.

Delete tasks: Remove tasks using todoguru delete taskID.

📦 Prerequisites

Go 1.16 or higher installed on your system.

Basic knowledge of Go and CLI applications.

🛠 Installation

Clone the repository:

git clone https://github.com/GuruOrGoru/ToDoGoru.git

Navigate to the project directory:

cd ToDoGoru

Build the project:

go build -o todoguru

This will create a todoguru binary in the current directory.

(Optional) Move the binary to your system's PATH to use it globally:

sudo mv todoguru /usr/local/bin/

📘 Usage

1. Add a Task

todoguru add "Task description here"

Example:

todoguru add "Buy groceries"

2. List All Tasks

todoguru list

This will display all tasks with their IDs, descriptions, and completion statuses.

3. Mark a Task as Complete

todoguru done <taskID>

Example:

todoguru done 1

4. Delete a Task

todoguru delete <taskID>

Example:

todoguru delete 2

📄 Example Session

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

🤝 Contributing

If you'd like to contribute to ToDoGoru, feel free to fork the repository and submit a pull request. We welcome bug fixes, feature improvements, and general enhancements.

Steps to Contribute:

Fork the repository

Create a new branch:

git checkout -b feature-name

Make your changes

Commit your changes:

git commit -am "Add new feature"

Push to the branch:

git push origin feature-name

Open a pull request

Happy task managing with ToDoGoru! 🧠📋

