package actions

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/sbszcz/clibook/repository"
	"github.com/sbszcz/clibook/ui"
	"github.com/urfave/cli/v3"
)

const (
	InvalidID = -1
)

type RootAction struct {
	DB *sql.DB
}

func (a *RootAction) Run(ctx context.Context, cmd *cli.Command) error {
	idString := cmd.String("id")
	id := int64(InvalidID)
	if len(idString) > 0 {

		parseResult, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			return fmt.Errorf("'%v' as value for '--id' not a valid (int64)", idString)
		}
		id = parseResult
	}

	notes, err := a.fetchNotes(id)
	if err != nil {
		return err
	}

	f := cmd.String("format")
	renderer := ui.NewRenderer(ui.Format(f))
	renderer.Render(notes)
	return nil
}

func (a *RootAction) fetchNotes(id int64) ([]repository.CliNote, error) {
	queries := repository.New(a.DB)

	if id == InvalidID {
		return queries.GetAll(context.Background())
	}

	note, err := queries.GetOne(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("Note with id '%v' not found", id)
	}

	return []repository.CliNote{note}, nil
}
