package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// if os.Args[1] == "s" {
	// 	m := Mask{mask: '*', space: " ", prefix: "http://"}
	// 	str := m.DisguiseStr("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")
	// 	//nolint:forbidigo
	// 	fmt.Println(str)
	// } else if os.Args[1] == "f" {
	// 	m := Mask{mask: '*', space: " ", prefix: "http://"}
	// 	err := m.DisguiseFile("links.txt", "disguised_links.txt")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:  "line",
				Usage: "line",
				Action: func(c *cli.Context) error {
					m := Mask{mask: '*', space: " ", prefix: "http://"}
					str := m.DisguiseStr("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")
					//nolint:forbidigo
					fmt.Println(str)
					return nil
				},
			},
			{
				Name:  "file",
				Usage: "file",
				Action: func(c *cli.Context) error {
					m := Mask{mask: '*', space: " ", prefix: "http://"}
					err := m.DisguiseFile("links.txt", "disguised_links.txt")
					if err != nil {
						panic(err)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
