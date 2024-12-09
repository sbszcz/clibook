package ui

import (
	"fmt"

	"github.com/sbszcz/clibook/repository"
)

type CsvRenderer struct{}

func (_ *CsvRenderer) Render(notes []repository.CliNote) {
	fmt.Println("ID,COMMAND,NOTE")
	for _, note := range notes {
		fmt.Printf("%v,%v,%v\n", note.ID, note.Command, note.Note)
	}
}
