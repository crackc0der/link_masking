package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:  "line",
				Usage: "line",
				Subcommands: []cli.Command{
					{
						Name:  "http",
						Usage: "http",
						Action: func(cCtx *cli.Context) error {
							m := Mask{prefix: "http://"}
							str := m.DisguiseStr("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")
							//nolint:forbidigo
							fmt.Println(str)
							return nil
						},
					},
					{
						Name:  "https",
						Usage: "https",
						Action: func(cCtx *cli.Context) error {
							m := Mask{prefix: "https://"}
							str := m.DisguiseStr("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")
							//nolint:forbidigo
							fmt.Println(str)
							return nil
						},
					},
				},
			},
			{
				Name:  "file",
				Usage: "file",
				Subcommands: []cli.Command{
					{
						Name:  "http",
						Usage: "http",
						Action: func(c *cli.Context) error {
							m := Mask{prefix: "http://"}
							err := m.DisguiseFile("links.txt", "disguised_links.txt")
							if err != nil {
								panic(err)
							}
							return nil
						},
					},
					{
						Name:  "https",
						Usage: "https",
						Action: func(c *cli.Context) error {
							m := Mask{prefix: "https://"}
							err := m.DisguiseFile("links.txt", "disguised_links.txt")
							if err != nil {
								panic(err)
							}
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
