package main

import "fmt"

func main() {
	var num0 int
	var num1 int = 1

	var num2 = 20
	fmt.Println(num0, num1, num2)

	num := 30

	num += 1

	num++
	fmt.Println(num)

	userIndex := 10
	user_index := 11
	fmt.Println(userIndex, user_index)

	var weigth, heigth = 10, 20

	weigth, age := 12, 22
	fmt.Println(weigth, heigth, age)
}
