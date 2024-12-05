package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/sbszcz/clibook/commands"
	"github.com/sbszcz/clibook/utils"
	"github.com/urfave/cli/v3"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func main() {
	dbDirectory := xdg.DataHome + "/clibook"
	dbFile := "clibook.db"

	dbFilePath, err := utils.EnsureFileExists(dbDirectory, dbFile)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite", dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(context.Background(), ddl); err != nil {
		log.Fatal(err)
	}

	tabRenderAction := commands.TabRenderAction{DB: db}
	addCommand := commands.AddCommand{DB: db}
	// exportCommand := commands.ExportCommand{DB: db}

	cmd := &cli.Command{
		Name:   "clibook",
		Usage:  "Maintain your command line diary directly from your favorite place.",
		Action: tabRenderAction.Render,
		Commands: []*cli.Command{
			addCommand.Create(),
			// exportCommand.Create(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
