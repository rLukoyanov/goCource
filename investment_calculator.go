package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(ratio)
}
