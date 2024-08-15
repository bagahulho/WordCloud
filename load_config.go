package main

import (
	"fmt"
	"github.com/flopp/go-findfont"
	"github.com/psykhi/wordclouds"
	"gopkg.in/yaml.v2"
	"image/color"
)

var conf = Conf{
	FontMaxSize:     300,
	FontMinSize:     30,
	RandomPlacement: false,
	Colors: []color.RGBA{
		{R: 247, G: 144, B: 30, A: 255},
		{R: 194, G: 69, B: 39, A: 255},
		{R: 38, G: 103, B: 118, A: 255},
		{R: 173, G: 210, B: 224, A: 255},
	},
	BackgroundColor: color.RGBA{R: 250, G: 250, B: 250, A: 255},
	Mask: MaskConf{"", color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}},
	Debug: false,
}

func loadConfig(pathToFile string, width, height int) ([]wordclouds.Option, error) {
	var fontPath string
	var errFont error
	fonts := []string{"arial.ttf", "Ubuntu-M.ttf", "FreeMono.ttf"}

	for _, font := range fonts {
		fontPath, errFont = findfont.Find(font)
		if errFont == nil {
			break
		}
	}
	if errFont != nil {
		return nil, fmt.Errorf("font not found in your system")
	}

	conf.FontFile = fontPath
	conf.Mask.File = pathToFile
	conf.Width = width
	conf.Height = height

	if conf.Debug {
		confYaml, _ := yaml.Marshal(conf)
		fmt.Printf("Configuration: %s\n", confYaml)
	}

	var boxes []*wordclouds.Box
	if conf.Mask.File != "" {
		boxes = wordclouds.Mask(
			conf.Mask.File,
			conf.Width,
			conf.Height,
			conf.Mask.Color)
	}

	colors := make([]color.Color, 0)
	for _, c := range conf.Colors {
		colors = append(colors, c)
	}

	optionsArr := []wordclouds.Option{wordclouds.FontFile(conf.FontFile),
		wordclouds.FontMaxSize(conf.FontMaxSize),
		wordclouds.FontMinSize(conf.FontMinSize),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(conf.Height),
		wordclouds.Width(conf.Width),
		wordclouds.RandomPlacement(conf.RandomPlacement),
		wordclouds.BackgroundColor(conf.BackgroundColor),
	}
	if conf.SizeFunction != nil {
		optionsArr = append(optionsArr, wordclouds.WordSizeFunction(*conf.SizeFunction))
	}
	if conf.Debug {
		optionsArr = append(optionsArr, wordclouds.Debug())
	}

	return optionsArr, nil
}
