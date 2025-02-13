package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var invsetmentAmount, years, expectedReturnRate float64

	fmt.Print("Invsetment Amount: ")
	fmt.Scan(&invsetmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := invsetmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
