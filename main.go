package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/sbszcz/clibook/ui/actions"
	"github.com/sbszcz/clibook/ui/commands"
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

	rootAction := actions.RootAction{DB: db}

	addNote := commands.AddNote{DB: db}

	cmd := &cli.Command{
		Name:  "clibook",
		Usage: "Maintain your command line diary directly from your favorite place.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Value:   "tab",
				Aliases: []string{"f"},
				Usage:   "Determines the output format of clibook content. Possible formats are: tab (default), csv",
			},
			&cli.StringFlag{
				Name:    "id",
				Value:   "",
				Aliases: []string{"i"},
				Usage:   "Only one item with given ID is selected.",
			},
		},
		Action: rootAction.Run,
		Commands: []*cli.Command{
			addNote.Create(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
