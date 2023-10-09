package main

import (
	"os"
)

func main() {
	if os.Args[1] == "s" {
		m := Mask{mask: '*', space: " ", prefix: "http://"}
		str := m.DisguiseStr("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")
		//nolint:forbidigo
		panic(str)
	} else if os.Args[1] == "f" {
		m := Mask{mask: '*', space: " ", prefix: "http://"}
		err := m.DisguiseFile("links.txt", "disguised_links.txt")
		if err != nil {
			panic(err)
		}
	}
}
