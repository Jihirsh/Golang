package main

import (
	"errors"
	"fmt"
)

func main() {
	var firstNumber int
	var operator string
	var secondNumber int

	fmt.Println("Enter first number:")
	_, err := fmt.Scan(&firstNumber)
	if err != nil {
		fmt.Println("Invalid operand")
		return
	}

	fmt.Println("Enter operator (+ - * /):")
	fmt.Scan(&operator)

	fmt.Println("Enter second number:")
	_, err = fmt.Scan(&secondNumber)
	if err != nil {
		fmt.Println("Invalid operand")
		return
	}

	switch operator {

	case "+":
		fmt.Println(add(firstNumber, secondNumber))

	case "-":
		fmt.Println(subtract(firstNumber, secondNumber))

	case "*":
		fmt.Println(multiply(firstNumber, secondNumber))

	case "/":
		result, err := divide(firstNumber, secondNumber)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)

	default:
		fmt.Println("Invalid Operator Choice! Please choose from the given operators.")
	}
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}

	return a / b, nil
}