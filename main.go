package main

import (
	"log"
	"os"

	"github.com/makesitgo/makes-text/commands"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tt",
		Usage: "do text transforms",
		Commands: []*cli.Command{
			&commands.Base64,
			&commands.JSON,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("makes-text failed: " + err.Error())
	}
}
