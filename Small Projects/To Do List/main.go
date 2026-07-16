package main

import (
	"os"
	"bufio"
	"fmt"
)


func main() {
	var option int = 0
	toDoList := []string{}

	for {
		option = showMenu()
		if option == 4 {break}
		manage(option, &toDoList)
	}
}

//exploring how to convert string to int

func showMenu() int {
	var localOption int

	fmt.Println()
	fmt.Println("==== TODO LIST ====")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Exit")
	fmt.Println()
	fmt.Print("Choose: ")

	fmt.Scan(&localOption)

	return localOption
}

func manage(option int, toDoList *[]string) {
	switch option {
	case 1:
		addTask(toDoList)
	case 2:
		viewTasks(*toDoList)
	case 3:
		*toDoList = deleteTask(*toDoList)
	default:
		fmt.Println("Choose from the given options")
	}
}


func addTask(toDoList *[]string) {
	fmt.Print("Enter task: ")
	task := bufio.NewScanner(os.Stdin)
	task.Scan()
	*toDoList = append(*toDoList, task.Text())
}

func viewTasks(toDoList []string) {
	for index, task := range toDoList {
		fmt.Println(index + 1, task)
	}
}

func deleteTask(toDoList []string) [] string {
	var index int

	if len(toDoList) == 0 {
		fmt.Println("No tasks to delete.")
		return toDoList
	}

	fmt.Print("Choose the index of the task you want to delete: ")
	fmt.Scan(&index)

	if index <= 0 || index > len(toDoList) {
		fmt.Println("Choose a valid index.")
		return toDoList
	}

	if index >= 1 {
		fmt.Printf("Removed the task %v from the ToDo List\n", toDoList[index - 1])
	}
	toDoList = append(toDoList[:index-1], toDoList[index:]...)

	return toDoList
}