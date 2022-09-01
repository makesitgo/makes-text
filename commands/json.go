package commands

import (
	"encoding/json"
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var (
	JSON = cli.Command{
		Name:  "json",
		Usage: "json text utilities",
		Subcommands: []*cli.Command{
			{
				Name:  "pretty",
				Usage: "pretty print the json payload",
				Flags: []cli.Flag{&cli.StringFlag{
					Name:     "value",
					Usage:    "the value to make pretty",
					Required: true,
				}},
				Action: prettyPrint,
			},
		},
	}
)

func prettyPrint(cliCtx *cli.Context) error {
	value := cliCtx.String("value")

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(value), &m); err != nil {
		return err
	}
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
