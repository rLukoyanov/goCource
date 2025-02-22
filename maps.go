package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	fmt.Println(len(userNames), cap(userNames))

	userNames = append(userNames, "Max")

	fmt.Println(userNames)

	coursesRatings := make(floatMap, 3)

	coursesRatings["go"] = 4.7
	coursesRatings["react"] = 4.8
	coursesRatings["angular"] = 4.7
	coursesRatings["node"] = 4.8

	coursesRatings.output()

	// fmt.Println(coursesRatings)

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range coursesRatings {
		fmt.Println(key, "_", value)
	}
}
