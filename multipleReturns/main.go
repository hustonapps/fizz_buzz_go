package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 3.5

func main() {
	var investmentAmount, expectedReturnRate, years float64;

	outputText("Investment Amount")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate")
	fmt.Scan(&expectedReturnRate)

	outputText("Years")
	fmt.Scan(&years)

	fv, rfv := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: $%.2f\n", fv)
	formattedRFV := fmt.Sprintf("Real Future Value: $%.2f\n", rfv)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText (text string) {
	fmt.Println(text);
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return
}