package commands

import (
	"context"
	"database/sql"

	"github.com/sbszcz/clibook/repository"
	"github.com/urfave/cli/v3"
)

const (
	AddCommandName  = "add"
	CommandFlagName = "command"
	NoteFlagName    = "note"
)

type AddNote struct {
	DB *sql.DB
}

func (c *AddNote) Create() *cli.Command {
	cmd := &cli.Command{
		Name:   AddCommandName,
		Usage:  "add a new note for a command",
		Action: c.handleAddCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  CommandFlagName,
				Value: "",
				Usage: "the actual cli command e.g. 'kubectl -n dev get pods'",
			},
			&cli.StringFlag{
				Name:  NoteFlagName,
				Value: "",
				Usage: "the note for the command",
			},
		},
	}
	return cmd
}

func (c *AddNote) handleAddCommand(ctx context.Context, cmd *cli.Command) error {
	command := cmd.String(CommandFlagName)
	if len(command) < 1 {
		return cli.ShowAppHelp(cmd)
	}
	note := cmd.String(NoteFlagName)
	if len(note) < 1 {
		return cli.ShowAppHelp(cmd)
	}

	q := repository.New(c.DB)
	q.CreateCliNote(ctx, repository.CreateCliNoteParams{
		Command: command,
		Note:    note,
	})

	return nil
}
