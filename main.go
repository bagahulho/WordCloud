package main

import (
	"fmt"
	"github.com/psykhi/wordclouds"
	"github.com/urfave/cli/v2"
	"image/png"
	"log"
	"os"
	"time"
)

var (
	pathToJson string
	pathToMask string
	outputPath string
	width      int
	height     int
)

func main() {
	app := cli.NewApp()
	app.Name = "wordcloud"
	app.Usage = "Makes a picture based on the number of words spoken in telegram"
	app.Description = "You can extract your correspondence from telegram and create a beautiful picture from your words. The more often a word occurs, the bigger it will be in the picture"
	app.Authors = []*cli.Author{
		{Name: "Baga", Email: "bagahulho06@gmail.com"},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "makeSingle",
			Aliases: []string{"mS"},
			Usage:   "Creating an image based on a single chat",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Destination: &pathToJson,
					Name:        "json",
					Aliases:     []string{"j"},
					Usage:       "Path to JSON file from Telegram",
					Required:    true,
				},
				&cli.StringFlag{
					Destination: &pathToMask,
					Name:        "mask",
					Aliases:     []string{"m"},
					Usage:       "Path to mask image (png) (picture without background)",
					Value:       "default",
				},
				&cli.StringFlag{
					Destination: &outputPath,
					Name:        "output",
					Aliases:     []string{"o"},
					Usage:       "Output path for the generated word cloud image",
					Value:       "wordcloud.png",
				},
				&cli.IntFlag{
					Destination: &width,
					Name:        "width",
					Aliases:     []string{"wi"},
					Usage:       "Width of the output image in pixels (Recommended: more than 2000)",
					Value:       2048,
				},
				&cli.IntFlag{
					Destination: &height,
					Name:        "height",
					Aliases:     []string{"he"},
					Usage:       "Height of the output image in pixels (Recommended: more than 2000)",
					Value:       2048,
				},
			},
			Action: func(c *cli.Context) error {
				return makeWordCloud(MakeMapSingle)
			},
		},
		{
			Name:    "makeMulti",
			Aliases: []string{"mM"},
			Usage:   "Creating an image based on all chats",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Destination: &pathToJson,
					Name:        "json",
					Aliases:     []string{"j"},
					Usage:       "Path to JSON file from Telegram",
					Required:    true,
				},
				&cli.StringFlag{
					Destination: &pathToMask,
					Name:        "mask",
					Aliases:     []string{"m"},
					Usage:       "Path to mask image (png) (picture without background)",
					Required:    true,
				},
				&cli.StringFlag{
					Destination: &outputPath,
					Name:        "output",
					Aliases:     []string{"o"},
					Usage:       "Output path for the generated word cloud image",
					Value:       "wordcloud.png",
				},
				&cli.IntFlag{
					Destination: &width,
					Name:        "width",
					Aliases:     []string{"wi"},
					Usage:       "Width of the output image in pixels (Recommended: more than 2000)",
					Value:       2048,
				},
				&cli.IntFlag{
					Destination: &height,
					Name:        "height",
					Aliases:     []string{"he"},
					Usage:       "Height of the output image in pixels (Recommended: more than 2000)",
					Value:       2048,
				},
			},
			Action: func(c *cli.Context) error {
				return makeWordCloud(MakeMapMulti)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func makeWordCloud(makeMap func(pathToJson string) (map[string]int, error)) error {
	fmt.Println("Starting to generate word cloud...")
	fmt.Printf("Using JSON file: %s\n", pathToJson)
	fmt.Printf("Using mask image: %s\n", pathToMask)
	fmt.Printf("Output will be saved to: %s\n", outputPath)
	fmt.Printf("Image size: %d x %d px\n", width, height)

	// Вызов основной логики для генерации word cloud
	words, err := makeMap(pathToJson)
	if err != nil {
		return err
	}
	draw(words, pathToMask, outputPath, width, height)

	fmt.Println("Word cloud generated successfully.")

	return nil
}

func draw(words map[string]int, pathToMask, outputPath string, width, height int) {
	start := time.Now()
	optionsArr, err := loadConfig(pathToMask, width, height)
	if err != nil {
		log.Fatal(err)
	}

	w := wordclouds.NewWordcloud(words,
		optionsArr...,
	)

	img := w.Draw()
	outputFile, err := os.Create(outputPath)
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
