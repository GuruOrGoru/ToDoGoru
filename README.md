## ToDoGoru - CLI Task Manager in Go ##
ToDoGoru is a simple command-line interface (CLI) task manager built in Go (Golang). It allows users to manage their tasks directly from the terminal by adding, listing, marking tasks as complete, and deleting them. This app is designed to be easy to use and efficient for managing everyday tasks.

Features
Add tasks: Create tasks by typing todoguru add "task".

Mark tasks as complete: Mark a task as done with todoguru done taskID.

List tasks: View all tasks with todoguru list.

Delete tasks: Remove tasks by using todoguru delete taskID.

Prerequisites
Go 1.16 or higher installed on your system.

Basic knowledge of Go and CLI applications.

Installation
Clone the repository:

bash
Copy
Edit
git clone https://github.com/GuruOrGoru/ToDoGoru.git
Navigate to the project directory:

bash
Copy
Edit
cd ToDoGoru
Build the project:

bash
Copy
Edit
go build -o todoguru
This will create a todoguru binary in the current directory.

Optionally, you can move the todoguru binary to a directory in your system's PATH to use it globally:

bash
Copy
Edit
sudo mv todoguru /usr/local/bin/
Usage
1. Add a Task
To add a task, use the following command:

bash
Copy
Edit
todoguru add "Task description here"
This will add the task to the list, and it will be assigned a unique task ID.

Example:

bash
Copy
Edit
todoguru add "Buy groceries"
2. List All Tasks
To view all tasks, run:

bash
Copy
Edit
todoguru list
This will display a list of all tasks with their IDs, descriptions, and completion statuses.

3. Mark a Task as Complete
To mark a task as complete, use the task's ID:

bash
Copy
Edit
todoguru done taskID
Example:

bash
Copy
Edit
todoguru done 1
This will mark the task with ID 1 as complete.

4. Delete a Task
To delete a task, use the following command with the task's ID:

bash
Copy
Edit
todoguru delete taskID
Example:

bash
Copy
Edit
todoguru delete 2
This will remove the task with ID 2 from the list.

Example
Here is an example of how the CLI application works:

bash
Copy
Edit
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
Contributing
If you'd like to contribute to ToDoGoru, feel free to fork the repository and submit a pull request. We welcome bug fixes, feature improvements, and general enhancements.

Fork the repository

Create a new branch (git checkout -b feature-name)

Make your changes

Commit your changes (git commit -am 'Add new feature')

Push to the branch (git push origin feature-name)

Open a pull request