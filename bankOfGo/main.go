package main

import "fmt"

var choiceMap = map[int]string {
	1: "Check balance",
	2: "Deposit money",
	3: "Withdraw money",
	4: "Exit",
}

func main() {
	var accountBalance float64 = 1000.0

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
		case 3:
			fmt.Print("Withdraw amount: $")
			var withdrawAmount float64;
			fmt.Scan(&withdrawAmount)
			if (withdrawAmount > accountBalance) {
				fmt.Println("\nInsufficient funds. Please try again.")
				continue
			}
			accountBalance -= withdrawAmount
			showBalance(accountBalance)
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