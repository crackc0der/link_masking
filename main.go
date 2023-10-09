package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "s" {
		m := Mask{mask: '*', space: " ", prefix: "http://"}
		str := m.DisguiseStr("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")
		//nolint:forbidigo
		fmt.Println(str)
	} else if os.Args[1] == "f" {

	}

}
