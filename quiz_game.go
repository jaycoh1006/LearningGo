package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Great! You're old enough to play.")
	} else {
		fmt.Println("Sorry! You're too young to play.")
		return
	}

	fmt.Printf("Who's the best NFL team? ")
	var answer string
	fmt.Scan(&answer)
	if answer != "Commanders" {
		fmt.Println("Wrong! The Washington Commanders are, in fact, the best team.")
	} else {
		fmt.Println("Correct! You clearly have superior taste.")
	}
}
