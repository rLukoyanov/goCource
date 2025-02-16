package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	getAdultYeasr(agePointer)
	fmt.Println("age", age)
}

func getAdultYeasr(age *int) {
	*age = *age - 18
}
