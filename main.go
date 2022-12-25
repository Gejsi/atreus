package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "atreus",
		Usage: "Avoid spam and tracking by using disposable data.",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g", "gen"},
				Usage:   "Generate some throwaway data",
				Action:  generate,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "Shows a list of messages sent to your temporary email",
				Action:  list,
				// TODO: add --recent and --id flags
				// Flags: []cli.Flag{
				// 	&cli.StringFlag{Name: "email", Aliases: []string{"e"}},
				// },
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
