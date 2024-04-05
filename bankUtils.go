package bankUtils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ShowMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println()
}

func ReadBalanceFromFile(filename string) (float64, error) {
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

func WriteBalanceToFile(balance float64, filename string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(filename, []byte(balanceText), 0644)
}