package ui

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sbszcz/clibook/repository"
)

type TabRenderer struct{}

func (_ *TabRenderer) Render(notes []repository.CliNote) {
	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(tabwriter, "ID\tCOMMAND\tNOTE")

	for _, note := range notes {
		fmt.Fprintf(tabwriter, "%v\t%s\t%s\n", note.ID, note.Command, note.Note)
	}

	tabwriter.Flush()
}
