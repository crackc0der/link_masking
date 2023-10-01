package main

import (
	"fmt"
	"strings"
)

func main() {
	str := disguise("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")
	//nolint:forbidigo
	fmt.Println(str)
}

const (
	mask   byte   = '*'       // disguise symbol
	space  string = " "       // line separator character
	prefix string = "http://" // a prefix that does not need to be masked
)

func disguise(str string) string {
	// if no occurrences are found, return false
	if !strings.Contains(str, prefix) {
		return str
	}

	// split the line into separate words
	words := strings.Fields(str) //
	finalArr := make([]string, 0, len(words))

	for _, word := range words {
		// if an occurrence is found
		if strings.Contains(word, prefix) {
			var strArr = []byte(word)

			// the first 7 elements of the word are "http://" they do not need to be masked
			for i := len(prefix); i != len(word); i++ {
				strArr[i] = mask // mask the link
			}

			// add a link to the final slice
			finalArr = append(finalArr, string(strArr))

			continue
		}
		// add the remaining words to the final slice
		finalArr = append(finalArr, word)
	}

	// we connect the elements of the slice into a string separated by a space and return the string
	return strings.Join(finalArr, space)
}
