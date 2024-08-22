package main

import "image/color"

type Conf struct {
	FontMaxSize     int    `yaml:"font_max_size"`
	FontMinSize     int    `yaml:"font_min_size"`
	RandomPlacement bool   `yaml:"random_placement"`
	FontFile        string `yaml:"font_file"`
	Colors          []color.RGBA
	BackgroundColor color.RGBA `yaml:"background_color"`
	Width           int
	Height          int
	Mask            MaskConf
	SizeFunction    *string `yaml:"size_function"`
	Debug           bool
}

type MaskConf struct {
	File  string
	Color color.RGBA
}

type SingleChat struct {
	Messages []struct {
		Type         string `json:"type"`
		TextEntities []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"text_entities"`
	} `json:"messages"`
}

type MultiChat struct {
	Chats struct {
		List []struct {
			Type     string `json:"type"`
			ID       int64  `json:"id"`
			Name     string `json:"name,omitempty"`
			Messages []struct {
				ID           int    `json:"id"`
				Type         string `json:"type"`
				TextEntities []struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"text_entities"`
			} `json:"messages"`
		} `json:"list"`
	} `json:"chats"`
}
