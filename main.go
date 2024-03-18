package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	balance float32
}

func intro(acc *Account) {

	fmt.Printf("welcome to your account, what do you want? \n 1. cash in \n 2. cash out \n 3. get balance \n")
	choice := ""

	fmt.Scanln(&choice)

	switch choice {
	case "1":
		cashIn(acc)
	case "2":
		cashOut(acc)
	case "3":
		getBalance(acc)
	default:
		fmt.Println("invalid choice (1, 2, 3)")
	}
}

func main() {

	var account Account

	for {

		intro(&account)
		quit := ""
		fmt.Println("(X) for quit, (Enter) for continue")
		fmt.Scanln(&quit)
		if quit == "x" {
			break
		}
	}
}

func getBalance(acc *Account) {
	fmt.Println("your current balance is:", acc.balance)
}

func cashIn(acc *Account) {
	var newBalance float32
	fmt.Println("How much money u wanna deposit? ")
	fmt.Scanln(&newBalance)

	if newBalance < 5 {
		fmt.Print("you can't desposit less than $5")
	} else {
		acc.balance += newBalance
		fmt.Println("your money was desposited successfully!")
		fmt.Println("your current balance is:", acc.balance)
	}
}

func cashOut(acc *Account) {
	var lessBalance float32

	fmt.Println("how much do you want to withdraw")
	fmt.Scanln(&lessBalance)

	if lessBalance < 2 {
		fmt.Println("you can't withdraw less than $2")
	} else {
		acc.balance -= lessBalance
	}
	fmt.Println("your current balance is:", acc.balance)
}
