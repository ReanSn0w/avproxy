package models

import (
	"errors"
)

var (
	ErrTitleSplit = errors.New("ошибка при разделении заголовка")
)

// Функция разбирает текст заголовка на 3 части, оригинальное название
func PrepareAVTitle(title string) (string, string, string) {
	iterator := 0
	res := []string{"", "", ""}

	for _, symbol := range title {
		if symbol == rune("/"[0]) || symbol == rune("["[0]) || symbol == rune("]"[0]) {
			iterator++
			continue
		}

		if iterator > 2 {
			break
		}

		res[iterator] = res[iterator] + string(symbol)
	}

	return clearString(res[0]), clearString(res[1]), clearString(res[2])
}

func clearString(input string) string {
	if input[0] == byte(" "[0]) {
		input = input[1:]
	}

	if input[len(input)-1] == byte(" "[0]) {
		input = input[:len(input)-1]
	}

	return input
}
