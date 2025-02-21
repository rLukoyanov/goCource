package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

// func main() {
// 	prices := [4]float64{1.23, 2, 43.34, 123.23}

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 123
// 	hiloghtetdPrices := featuredPrices[:1]
// 	fmt.Println(len(hiloghtetdPrices), cap(hiloghtetdPrices))

// 	hiloghtetdPrices = hiloghtetdPrices[:3]
// 	fmt.Println(len(hiloghtetdPrices), cap(hiloghtetdPrices))
// }

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(len(prices), cap(prices), prices)

	newPrices := append(prices, 5.99)
	fmt.Println(len(newPrices), cap(newPrices), prices)

	discountPrices := []float64{10.99, 8.99}
	prices = append(prices, discountPrices...)
}
