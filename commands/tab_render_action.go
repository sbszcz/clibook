package commands

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sbszcz/clibook/repository"
	"github.com/urfave/cli/v3"
)

type TabRenderAction struct {
	DB *sql.DB
}

func (a *TabRenderAction) Render(ctx context.Context, cmd *cli.Command) error {
	queries := repository.New(a.DB)
	cliNotes, err := queries.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch all cli notes: %w", err)
	}

	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(tabwriter, "COMMAND\tNOTE")

	for _, note := range cliNotes {
		fmt.Fprintf(tabwriter, "%s\t%s\n", note.Command, note.Note)
	}

	tabwriter.Flush()

	return nil
}
