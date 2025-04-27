package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const PATH = "notes.json"

type note struct {
	Title   string
	Content string
	Created string // Changed to string
}

func (n note) outputNoteJson() {
	b, err := json.MarshalIndent(n, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}

	os.WriteFile(PATH, b, 0755)
}

func main() {
	title := getUserData("Please input title: ")
	content := getUserData("Please input content: ")

	var appNote note
	appNote = note{
		Title:   title,
		Content: content,
		Created: time.Now().Format(time.DateTime),
	}

	appNote.outputNoteJson()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
