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
