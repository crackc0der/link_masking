package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type Mask struct {
	prefix string
}

func (m *Mask) DisguiseStr(str string) string {
	space := " "

	// split the line into separate words
	words := strings.Fields(str) //

	finalArr := m.masking(words)

	// we connect the elements of the slice into a string separated by a space and return the string
	return strings.Join(finalArr, space)
}

func (m *Mask) DisguiseFile(path string, disguisedLinks string) error {
	var data []byte

	space := " "

	file, errRead := ioutil.ReadFile(path)
	if errRead != nil {
		panic(errRead)
	}
	// break all the lines into words
	data = append(data, file...)

	// convert the byte array to a string array

	words := strings.Fields(string(data))

	f, err := os.OpenFile(disguisedLinks, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	finalArr := strings.Join(m.masking(words), space)

	// for _, word := range finalArr {
	_, errOpen := f.WriteString(finalArr)
	if errOpen != nil {
		panic(errOpen)
	}
	// }

	return nil
}

func (m *Mask) masking(words []string) []string {
	finalArr := make([]string, 0, len(words))

	var mask byte = 42

	prefixHTTP := "http"
	prefixHTTPS := "https"
	sufixProtocolLen := 3

	for _, word := range words {
		// if an occurrence is found
		if strings.Contains(word, prefixHTTPS) {
			var strArr = []byte(word)
			// the first 7 elements of the word are "http://" they do not need to be masked
			for i := len(prefixHTTPS) + sufixProtocolLen; i != len(word); i++ {
				strArr[i] = mask // mask the link
			}

			// add a link to the final slice
			finalArr = append(finalArr, string(strArr))

			continue
		} else if strings.Contains(word, prefixHTTP) {
			var strArr = []byte(word)
			// the first 7 elements of the word are "http://" they do not need to be masked
			for i := len(prefixHTTP) + sufixProtocolLen; i != len(word); i++ {
				strArr[i] = mask // mask the link
			}
			// add a link to the final slice
			finalArr = append(finalArr, string(strArr))

			continue
		}
		// add the remaining words to the final slice
		finalArr = append(finalArr, word)
	}

	return finalArr
}
