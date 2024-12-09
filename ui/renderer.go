package ui

import "github.com/sbszcz/clibook/repository"

type Format string

type Renderer interface {
	Render(notes []repository.CliNote)
}

func NewRenderer(f Format) Renderer {
	switch f {
	case "csv":
		return &CsvRenderer{}
	case "tab":
		return &TabRenderer{}
	default:
		return &TabRenderer{}
	}
}
