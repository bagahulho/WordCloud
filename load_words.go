package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonChat struct {
	Name     string    `json:"name"`
	ID       int64     `json:"id"`
	Messages []Message `json:"messages"`
}

type Message struct {
	ID           int          `json:"id"`
	Date         string       `json:"date"`
	From         string       `json:"from,omitempty"`
	FromID       string       `json:"from_id,omitempty"`
	Type         string       `json:"type"`
	TextEntities []TextEntity `json:"text_entities"`
}

type TextEntity struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func MakeMap(pathToFile string) (map[string]int, error) {
	fileData, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	data := JsonChat{}
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling file: %w", err)
	}
	result := make(map[string]int)
	for _, message := range data.Messages {
		if len(message.TextEntities) == 0 || message.TextEntities[0].Type != "plain" {
			continue
		}
		words := ExtractWords(message.TextEntities[0].Text)
		for _, word := range words {
			if len([]rune(word)) > 4 {
				word = strings.ToLower(word)
				result[word] += 1
			}
		}
	}
	return result, nil
}
