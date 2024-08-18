package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func MakeMapSingle(pathToFile string) (map[string]int, error) {
	fileData, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
  
	data := SingleChat{}
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

func MakeMapMulti(pathToFile string) (map[string]int, error) {
	fileData, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	data := MultiChat{}
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling file: %w", err)
	}

	result := make(map[string]int)
	for _, chat := range data.Chats.List {
		for _, message := range chat.Messages {
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
	}

	return result, nil
}
