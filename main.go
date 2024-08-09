package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"image/color"
	"image/png"
	"log"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/psykhi/wordclouds"
)

var config = flag.String("config", "config.yaml", "path to config file")
var output = flag.String("output", "output.png", "path to output image")

func main() {
	// Load words
	var pathToFile string
	flag.StringVar(&pathToFile, "path", "", "path to input file")
	flag.Parse()
	result, err := MakeMap(pathToFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load config
	var conf Conf
	content, err := os.ReadFile(*config)
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

	start := time.Now()
	oarr := []wordclouds.Option{wordclouds.FontFile(conf.FontFile),
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
		oarr = append(oarr, wordclouds.WordSizeFunction(*conf.SizeFunction))
	}
	if conf.Debug {
		oarr = append(oarr, wordclouds.Debug())
	}
	w := wordclouds.NewWordcloud(result,
		oarr...,
	)

	img := w.Draw()
	outputFile, err := os.Create(*output)
	if err != nil {
		panic(err)
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, img)

	// Don't forget to close files
	outputFile.Close()
	fmt.Printf("Done in %v\n", time.Since(start))
}
