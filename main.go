package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID        int
	TaskToDo  string
	Completed bool
	CreatedAt time.Time
}

var rootCmd = &cobra.Command{
	Use:   "mediaguru",
	Short: "A CLI tool for managing media files",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the database",
	Args:  cobra.ExactArgs(1),
	Run:   addTask,
}

var readAllCmd = &cobra.Command{
	Use:   "list",
	Short: "Read all tasks from the database",
	Run:   readAllTasks,
}

var completeCmd = &cobra.Command{
	Use:   "done",
	Short: "Complete a task",
	Args:  cobra.ExactArgs(1),
	Run:   completeTask,
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run:   deleteTask,
}

var tasks []Task = loadTasks()

func main() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(readAllCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addTask(cmd *cobra.Command, args []string) {
	file, err := os.OpenFile("task.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	task := Task{
		ID:        len(tasks) + 1,
		TaskToDo:  args[0],
		Completed: false,
		CreatedAt: time.Now(),
	}
	writer.Write([]string{fmt.Sprintf("%d", task.ID), task.TaskToDo, fmt.Sprintf("%t", task.Completed), task.CreatedAt.Format("2006-01-02 15:04:05")})
	tasks = append(tasks, task)
	fmt.Printf("Task added successfully: %s\n", task.TaskToDo)
}

func readAllTasks(cmd *cobra.Command, args []string) {
	fmt.Println("ID   | Task Description          | Completed | Created At")
	fmt.Println("--------------------------------------------------------")
	for _, task := range tasks {
		completedStr := "No"
		if task.Completed {
			completedStr = "Yes"
		}
		fmt.Printf("%-5d | %-25s | %-9s | %-20s\n", task.ID, task.TaskToDo, completedStr, task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func csvData(filePath string) (record [][]string, err error) {
	csvFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records, nil
}

func completeTask(cmd *cobra.Command, args []string) {
	file, err := os.OpenFile("task.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	id, err := ConvertStringToInt(args[0])
	if err != nil {
		fmt.Println("Error converting ID:", err)
		return
	}
	taskFound := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Printf("Task %d marked as completed\n", id)
			taskFound = true
			break
		}
	}
	if !taskFound {
		fmt.Printf("Task %d not found\n", id)
		return
	}
	for _, task := range tasks {
		writer.Write([]string{
			fmt.Sprintf("%d", task.ID),
			task.TaskToDo,
			fmt.Sprintf("%t", task.Completed),
			task.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
}

func loadTasks() []Task {
	var loadedTasks []Task
	records, err := csvData("task.csv")
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		id, err := ConvertStringToInt(record[0])
		if err != nil {
			fmt.Println("Error converting ID:", err)
			return loadedTasks
		}

		completed, err := ConvertStringToBool(record[2])
		if err != nil {
			fmt.Println("Error converting Completed:", err)
			return loadedTasks
		}

		createdAt, err := ConvertStringToTime(record[3])
		if err != nil {
			fmt.Println("Error converting CreatedAt:", err)
			return loadedTasks
		}
		task := Task{
			ID:        id,
			TaskToDo:  record[1],
			Completed: completed,
			CreatedAt: createdAt,
		}
		loadedTasks = append(loadedTasks, task)
	}
	return loadedTasks
}

func deleteTask(cmd *cobra.Command, args []string) {
	id, err := ConvertStringToInt(args[0])
	if err != nil {
		fmt.Println("Error converting ID:", err)
		return
	}
	taskFound := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = slices.Delete(tasks, i, i+1)
			fmt.Printf("Task %d deleted successfully\n", id)
			taskFound = true
			for j := range tasks {
				tasks[j].ID = j + 1
			}
			break
		}
	}
	if !taskFound {
		fmt.Printf("Task %d not found\n", id)
		return
	}
	file, err := os.OpenFile("task.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, task := range tasks {
		writer.Write([]string{
			fmt.Sprintf("%d", task.ID),
			task.TaskToDo,
			fmt.Sprintf("%t", task.Completed),
			task.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
}
