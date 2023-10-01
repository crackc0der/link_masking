package main

import (
	"fmt"
	"strings"
)

func main() {
	str, _ := disguise("1 http://11111111111 2 http://22222222222 http://sdfaiojdsfoi http://sd8923892348923")

	//nolint:forbidigo
	fmt.Println(str)
}

func disguise(str string) (string, bool) {
	var (
		strArr      []byte
		finalString []string
		words       []string
	)

	const (
		mask  byte   = '*' // символ маскировки
		space string = " " // символ разделения слов в строке
	)

	if !strings.Contains(str, "http://") { // если не найдено ни одного вхождения возвращаем false
		return "", false
	}

	words = strings.Fields(str) // разбиваем строку на отдельные слова

	for _, word := range words {
		if strings.Contains(word, "http://") { // если найдено вхождение
			for i := 0; i != len(word); i++ {
				strArr = append(strArr, word[i]) // добавляем его в промежуточный слайс
			}
			for i := 7; i != len(word); i++ { // первые 7 элементов слова это "http://" их маскировать ненадо
				strArr[i] = mask // маскируем ссылку
			}

			finalString = append(finalString, string(strArr)) // добавляем ссылку в финальный слайс
			strArr = nil
		} else {
			finalString = append(finalString, word) // добавляем остальные слова в финальный слайс
		}
	}

	return strings.Join(finalString, space), true // соединяем элементы слайса в строку через пробел и возвращаем строку
}
