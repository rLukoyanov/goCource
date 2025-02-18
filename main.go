package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exmple.com/main/note"
	"exmple.com/main/todo"
)

type Displayer interface {
	Display()
}

type Saver interface {
	Save() error
}

type outputtable interface {
	Saver
	Displayer
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething("123")
	title, content := getNoteData()
	textTodo := getUserInput("Todo text: ")

	userNote, err := note.New(title, content)
	if err != nil {
		panic(err)
	}

	todo, err := todo.New(textTodo)

	if err != nil {
		panic(err)
	}

	outputData(userNote)
	outputData(todo)
}

func printSomething(value any) {
	typedVal, ok := value.(int)
	switch value.(type) {
	case int:
		fmt.Println("integer")
	case float64:
		fmt.Println("float")
	default:
		fmt.Println(value)
	}

}

func outputData(data outputtable) {
	data.Display()
	saveData(data)
}

func saveData(data Saver) {
	err := data.Save()

	if err != nil {
		panic(err)
	}

	fmt.Println("saving succesed")
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.Trim(text, "\n")
	text = strings.Trim(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Title note: ")
	content := getUserInput("Content note: ")

	return title, content
}
