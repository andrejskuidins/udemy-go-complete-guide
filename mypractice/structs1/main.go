package main

import (
	"fmt"
	"time"
	"encoding/json"
	"os"
)

type note struct {
	title: string
	content: string
	created: time.Now()
}

func (n note) outputNoteJson() {
	// ...

	err := os.WriteFile("note.json", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func main()  {
	title := getUserData("Please input title: ")
	content := getUserData("Please input content: ")

	var appNote note
	appNote = note{
		title: title
		content: content
		created: created
	}

	appNote.outputNoteJson()
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}