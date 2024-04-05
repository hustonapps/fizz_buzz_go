package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const filename = "Balance.txt"

var choiceMap = map[int]string {
	1: "Check balance",
	2: "Deposit money",
	3: "Withdraw money",
	4: "Exit",
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(filename, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	balanceText, err := os.ReadFile(filename)
	if err != nil {
		return 0, errors.New("failed to read file")
	}

	balance, err := strconv.ParseFloat(string(balanceText), 64)
	if err != nil {
		return 0, errors.New("failed to parse balance")
	}

	return balance, nil
}

func main() {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("---------")
		// panic() shows more of a stack trace
	}

	fmt.Println("\nWelcome to Go Bank!")

	for {
		showMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showBalance(accountBalance)
		case 2:
			fmt.Print("Deposit amount: $")
			var depositAmount float64;
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("\nInvalid deposit amount. Please try again.")
				continue
			}

			accountBalance += depositAmount
			showBalance(accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdraw amount: $")
			var withdrawAmount float64;
			fmt.Scan(&withdrawAmount)
			if (withdrawAmount > accountBalance) {
				fmt.Println("\nInsufficient funds. Please try again.")
				continue
			} else if (withdrawAmount <= 0) {
				fmt.Println("\nInvalid withdrawal amount. Please try again.")
				continue
			}
			accountBalance -= withdrawAmount
			showBalance(accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("\nThank you for using Go Bank!")
			return
		}
	}
}

func showMenu() {
	fmt.Println("What do you want to do?")
	for k, v := range choiceMap {
		fmt.Printf("%d. %s\n", k, v)
	}
	fmt.Println()
}

func showBalance(bal float64) {
	fmt.Printf("\nYour account balance is $%.2f\n\n", bal)
}