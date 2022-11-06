package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App {
    Commands: []*cli.Command {
      {
        Name: "generate",
        Aliases: []string{"g", "gen"},
        Usage: "Generate some throwaway data.",
        Action: func(ctx *cli.Context) error {
          fmt.Println("generated data")
          return nil
        },
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
