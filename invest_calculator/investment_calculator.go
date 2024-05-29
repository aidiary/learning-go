package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Enter investment amount: ")
	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)	
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
}
