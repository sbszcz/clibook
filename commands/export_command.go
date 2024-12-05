package commands

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/urfave/cli/v3"
)

const (
	ExportCommandName = "export"
	FormatFlagName    = "format"
)

type ExportCommand struct {
	DB *sql.DB
}

func (c *ExportCommand) Create() *cli.Command {
	cmd := &cli.Command{
		Name:   ExportCommandName,
		Usage:  "exports cli notebook into a available format (markdown)",
		Action: c.handleExportCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  FormatFlagName,
				Value: "markdown",
				Usage: "format for export (currently only markdown)",
			},
		},
	}
	return cmd
}

func (c *ExportCommand) handleExportCommand(ctx context.Context, cmd *cli.Command) error {
	if len(cmd.String(FormatFlagName)) < 1 {
		cli.ShowAppHelp(cmd)
	}

	// TODO: render notenook entries as markdown definition list to stdout
	fmt.Println("export all the things")

	return nil
}
