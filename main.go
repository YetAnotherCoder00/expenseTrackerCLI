package main

import (
	"fmt"
)

var balance float32

func main() {

	fmt.Printf("welcome to your account, what do you want? \n 1. cash in \n 2. cash out \n 3. get balance \n")
	choice := ""

	fmt.Scanln(&choice)

	switch choice {
	case "1":
		cashIn()
	case "2":
		cashOut()
	case "3":
		getBalance()
	default:
		fmt.Println("invalid choice (1, 2, 3)")
	}

}

func getBalance() float32 {
	return balance
}

func cashIn() {

	var newBalance float32

	fmt.Println("How much money u wanna deposit? ")
	fmt.Scanln(&newBalance)

	if newBalance < 5 {
		fmt.Print("you can't desposit less than $5")
	} else {
		balance += newBalance
		fmt.Println("your money was desposited successfully!")
		fmt.Println("your current balance is:", newBalance)
	}
}

func cashOut() {
	var lessBalance float32

	fmt.Println("how much do you want to withdraw")
	fmt.Scanln(&lessBalance)

	if lessBalance < 2 {
		fmt.Println("you can't withdraw less than $2")
	} else {
		balance -= lessBalance
	}
}
