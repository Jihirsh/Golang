package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func showMenu() int {
	fmt.Println()
	fmt.Println("==== TODO LIST ====")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Exit")
	fmt.Println()
	fmt.Print("Choose: ")

	option:= readInt()
	return option
}

// Orchestrator function that calls other functions to do task based on the menu option chosen by the user
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
	value, readerr := readLine()
	if readerr != nil {
		fmt.Println(readerr)
	}

	*toDoList = append(*toDoList, value)
}

func viewTasks(toDoList []string) {
	for index, task := range toDoList {
		fmt.Println(index + 1, task)
	}
}

func deleteTask(toDoList []string) [] string {

	// If toDoList is empty, we do early exit
	if len(toDoList) == 0 {
		fmt.Println("No tasks to delete.")
		return toDoList
	}

	// Taking input from user in terminal and converting the string index type to int
	fmt.Print("Choose the index of the task you want to delete: ")
	index := readInt()

	// Validation check
	if index <= 0 || index > len(toDoList) {
		fmt.Println("Choose a valid index.")
		return toDoList
	}

	// Removing the task from the slice
	fmt.Printf("Removed the task %v from the ToDo List\n", toDoList[index - 1])
	toDoList = append(toDoList[:index-1], toDoList[index:]...)

	return toDoList
}

// Taking user input at different instances
func readLine() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {

		if scanner.Err() != nil {
			return "", scanner.Err()
		}
	}

	value := scanner.Text()

	return value, nil
}


// Makes sure valid input is enterred
func readInt () int {

	value, readerr := readLine()

	for readerr != nil {
		fmt.Println(readerr)
		fmt.Print("Please enter again: ")
		value, readerr = readLine()
	}

	option, err := strconv.Atoi(value)
	for err != nil {
		fmt.Println(err)
		fmt.Print("Choose a valid index: ")
		value, readerr = readLine()
		option, err = strconv.Atoi(value)
	}

	return option
}