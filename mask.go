package main

import (
	"fmt"
	"os"
	"strings"
)

type Mask struct {
	mask   byte
	space  string
	prefix string
}

func (m *Mask) DisguiseStr(str string) string {
	// if no occurrences are found, return false
	if !strings.Contains(str, m.prefix) {
		return str
	}

	// split the line into separate words
	words := strings.Fields(str) //
	finalArr := make([]string, 0, len(words))

	for _, word := range words {
		// if an occurrence is found
		if strings.Contains(word, m.prefix) {
			var strArr = []byte(word)

			// the first 7 elements of the word are "http://" they do not need to be masked
			for i := len(m.prefix); i != len(word); i++ {
				strArr[i] = m.mask // mask the link
			}

			// add a link to the final slice
			finalArr = append(finalArr, string(strArr))

			continue
		}
		// add the remaining words to the final slice
		finalArr = append(finalArr, word)
	}

	// we connect the elements of the slice into a string separated by a space and return the string
	return strings.Join(finalArr, m.space)
}

func (m *Mask) DisguiseFile(path string) bool {
	file, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	for _, str := range file {
		fmt.Println(str)
	}
	return true
}

func NewMask() {

}
