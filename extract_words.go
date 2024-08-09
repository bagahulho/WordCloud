package main

import "regexp"

var reg = regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ0-9]+(?:['-][a-zA-Zа-яА-ЯёЁ0-9]+)*`)

func ExtractWords(text string) []string {
	// Регулярное выражение для поиска слов, включая слова с дефисами и апострофами
	words := reg.FindAllString(text, -1)

	return words
}
