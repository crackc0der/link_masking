package main

import (
	"os"
)

type Mask struct {
}

func (m *Mask) DisguiseStr(str string) string {
	return m.masking(str)
}

func (m *Mask) DisguiseFile(path string, disguisedLinks string) error {
	file, errRead := os.ReadFile(path)
	if errRead != nil {
		return errRead
	}

	data := m.masking(string(file))

	endFile, err := os.OpenFile(disguisedLinks, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	_, errOpen := endFile.WriteString(data)
	if errOpen != nil {
		return errOpen
	}

	return nil
}

func (m *Mask) masking(str string) string {
	var mask byte = 42

	prefixHTTP := "http://"
	prefixHTTPS := "https://"
	var strArr = []byte(str)

	var space byte = 32

	http := m.KMPSearch(str, prefixHTTP)
	https := m.KMPSearch(str, prefixHTTPS)

	for _, v := range http {
		for i := v + len(prefixHTTP); i != len(str); i++ {
			if strArr[i] == space {
				break
			}

			strArr[i] = mask
		}
	}

	for _, v := range https {
		for i := v + len(prefixHTTPS); i != len(str); i++ {
			if strArr[i] == space {
				break
			}

			strArr[i] = mask
		}
	}

	return string(strArr)
}

func (m *Mask) KMPSearch(text, pattern string) []int {
	occurrences := []int{}
	textSize := len(text)
	patternSize := len(pattern)
	lps := m.computeLPSArray(pattern, patternSize)

	i := 0
	j := 0

	for i < textSize {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == patternSize {
			occurrences = append(occurrences, i-j)
			j = lps[j-1]
		} else if i < textSize && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return occurrences
}

func (m *Mask) computeLPSArray(pattern string, a int) []int {
	lps := make([]int, a)
	length := 0
	i := 1

	for i < a {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}
