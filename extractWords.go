package main

import "regexp"

func ExtractWords(text string) []string {
	// Регулярное выражение для поиска слов, включая слова с дефисами и апострофами
	reg := regexp.MustCompile(`[a-zA-Zа-яА-ЯёЁ0-9]+(?:['-][a-zA-Zа-яА-ЯёЁ0-9]+)*`)
	words := reg.FindAllString(text, -1)

	return words
}
