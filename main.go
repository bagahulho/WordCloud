package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	_ "net/http/pprof"
	"os"
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

	start := time.Now()
	optionsArr := loadConfig(*config)

	w := wordclouds.NewWordcloud(result,
		optionsArr...,
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
