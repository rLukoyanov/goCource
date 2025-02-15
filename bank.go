package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposite money")
		fmt.Println("3. Widraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)

		case 2:
			fmt.Println("Your deposit: ", accountBalance)
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount& Must be grater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("widrawl Amount: ", accountBalance)
			var widrawlAmount float64
			fmt.Scan(&widrawlAmount)

			if widrawlAmount <= 0 {
				fmt.Println("Invalid amount Must be grater than 0")
				continue
			}

			if widrawlAmount > accountBalance {
				fmt.Println("invalid amount. you cant withdraw more than you have!")
				continue
			}

			accountBalance -= widrawlAmount
			fmt.Println("Balance updated! New amount is:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("thanks for choosing our bank")
			fmt.Print("Goodbye!")
			return
		}

	}

}
