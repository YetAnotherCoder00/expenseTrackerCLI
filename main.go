package main

import (
	"fmt"
	"log"
	"strconv"
)

var balance float32

func main() {
	choice := ""

	const (
		cashIn  = 1
		cashOut = 2
		balance = 3
	)

	f, _ := strconv.ParseFloat(cashIn(), 64)

	switch choice {
	case cashIn:
		cashIn()

	}

}

func getBalance() float32 {
	return balance
}

func cashIn(newBalance float32) (string, error) {
	fmt.Println("How much money u wanna deposit? ")
	fmt.Scanln(&newBalance)

	if newBalance < 5 {
		fmt.Print("you can't desposit less that $4")
	}
	balance += newBalance

	if err != nil {
		log.Fatal("something went wrong")
	}

	f, _ := strconv.ParseFloat(newBalanc, 32)

	message := "your money was desposited successfully!"

	return err / message
}
