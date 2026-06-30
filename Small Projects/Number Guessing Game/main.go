package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(100) + 1
	var guess int

	// fmt.Printf("The random number that got generated is %v\n",random)

	fmt.Println("Guess a number between 1 and 100: ")
	fmt.Scan(&guess)

	for attempts := 10; attempts > 0; attempts-- {
		if random == guess {
			fmt.Println("You Win!")
			break
		} else if attempts > 1 {
			if random > guess {
				fmt.Println("Too Low")
				fmt.Printf("Attempts remaining: %v\n", attempts - 1)
				fmt.Scan(&guess)
			} else if random < guess {
				fmt.Println("Too High")
				fmt.Printf("Attempts remaining: %v\n", attempts - 1)
				fmt.Scan(&guess)
			}
		} else {
			fmt.Println("You Lose")
			fmt.Printf("The number was %v", random)
		}
	}
}