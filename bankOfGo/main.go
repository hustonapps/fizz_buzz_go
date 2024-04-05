package main

import (
	bankUtils "fizz_buzz_go"
	"fmt"
)

const filename = "Balance.txt"

func main() {
	accountBalance, err := bankUtils.ReadBalanceFromFile(filename)

	if err != nil {
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("---------")
		// panic() shows more of a stack trace
	}

	fmt.Println("\nWelcome to Go Bank!")

	for {
		bankUtils.ShowMenu()

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
			bankUtils.WriteBalanceToFile(accountBalance, filename)
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
			bankUtils.WriteBalanceToFile(accountBalance, filename)
		default:
			fmt.Println("\nThank you for using Go Bank!")
			return
		}
	}
}

func showBalance(bal float64) {
	fmt.Printf("\nYour account balance is $%.2f\n\n", bal)
}