package main

import (
	"fmt"
	"github.com/psykhi/wordclouds"
	"gopkg.in/yaml.v2"
	"image/color"
	"os"
	"path/filepath"
)

func loadConfig(pathToFile string) []wordclouds.Option {
	var conf Conf
	content, err := os.ReadFile(pathToFile)
	if err == nil {
		err = yaml.Unmarshal(content, &conf)
		if err != nil {
			fmt.Printf("Failed to decode config, using defaults instead: %s\n", err)
		}
	} else {
		fmt.Println("No config file. Using defaults")
	}
	os.Chdir(filepath.Dir(*config))

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

	return optionsArr
}
