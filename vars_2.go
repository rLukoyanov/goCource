package main

import "fmt"

func main() {
	var i int // платформозависмый
	var auto = -10

	var bigInt int64 = 1<<32 - 1
	var unInt uint64 = 12312
	var unBigInt uint64 = 1<<32 - 1

	var pi float32 = 3.141
	var e float64 = 2.718
	autoFloat := 1.232

	fmt.Println(pi, e, autoFloat)

	var b bool
	var isOk bool = true
	var suc = true
	c := false

	var com complex128 = -1.1 + 7.12i
}
