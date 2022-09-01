package commands

import (
	"encoding/base64"
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var (
	Base64 = cli.Command{
		Name:  "base64",
		Usage: "base64 text utilities",
		Subcommands: []*cli.Command{
			{
				Name:  "encode",
				Usage: "encode string to base64 representation",
				Flags: []cli.Flag{&cli.StringFlag{
					Name:     "value",
					Usage:    "the value to encode",
					Required: true,
				}},
				Action: base64Encode,
			},
			{
				Name:  "decode",
				Usage: "decode string from base64 representation",
				Flags: []cli.Flag{&cli.StringFlag{
					Name:     "value",
					Usage:    "the value to decode",
					Required: true,
				}},
				Action: base64Decode,
			},
		},
	}
)

func base64Decode(cliCtx *cli.Context) error {
	value := cliCtx.String("value")
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func base64Encode(cliCtx *cli.Context) error {
	value := cliCtx.String("value")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(value)))
	return nil
}
