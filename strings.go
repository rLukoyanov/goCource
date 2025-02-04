package main

import "unicode/utf8"

func main() {
	var str string

	var hello string = "Привет\n\t"
	var wor string = `Привет\n\t`
	var raw byte = '\x27'

	concatinate := hello + "Хеллоу"

	byteLen := len(hello)
	symbols := utf8.RuneCountInString(hello)

	byteString := []byte(hello)
	world := string(byteString)
}
