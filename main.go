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
	_ "github.com/urfave/cli/v2"
)

var output = flag.String("output", "output.png", "path to output image")

func main() {
	// Load words
	var pathToJson string
	var pathToMask string
	flag.StringVar(&pathToJson, "json", "", "path to json file")
	flag.StringVar(&pathToMask, "mask", "", "path to mask image")
	flag.Parse()
	result, err := MakeMap(pathToJson)
	if err != nil {
		log.Fatal(err)
	}

	// Load config

	start := time.Now()
	optionsArr := loadConfig(pathToMask)

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
