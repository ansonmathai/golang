package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    tasks := []string{}

    for {
        fmt.Println("1. Add Task")
        fmt.Println("2. List Tasks")
        fmt.Println("3. Mark Task as Done")
        fmt.Println("4. Delete Task")
        fmt.Println("5. Exit")
        fmt.Print("Enter your choice: ")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            fmt.Print("Enter task: ")
            task, _ := reader.ReadString('\n')
            tasks = append(tasks, strings.TrimSpace(task))
        case "2":
            fmt.Println("Tasks:")
            for i, task := range tasks {
                fmt.Printf("%d. %s\n", i+1, task)
            }
        case "3":
            fmt.Print("Enter task number to mark as done: ")
            taskNumStr, _ := reader.ReadString('\n')
            taskNumStr = strings.TrimSpace(taskNumStr)
            taskNum, err := strconv.Atoi(taskNumStr)
            if err == nil && taskNum > 0 && taskNum <= len(tasks) {
                tasks[taskNum-1] = tasks[taskNum-1] + " (done)"
            } else {
                fmt.Println("Invalid task number")
            }
        case "4":
            fmt.Print("Enter task number to delete: ")
            taskNumStr, _ := reader.ReadString('\n')
            taskNumStr = strings.TrimSpace(taskNumStr)
            taskNum, err := strconv.Atoi(taskNumStr)
            if err == nil && taskNum > 0 && taskNum <= len(tasks) {
                tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
            } else {
                fmt.Println("Invalid task number")
            }
        case "5":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}
