package main

import (
	"embed"
	"github.com/psykhi/wordclouds"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

//go:embed example/mask.png
var defaultMask embed.FS

func myMask(path string, width int, height int, exclude color.RGBA) []*wordclouds.Box {
	res := make([]*wordclouds.Box, 0)
	var img image.Image

	if path == "default" {
		file, err := defaultMask.Open("example/mask.png")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		img, err = png.Decode(file)
	} else {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		img, err = png.Decode(file)
	}

	// scale
	imgw := img.Bounds().Dx()
	imgh := img.Bounds().Dy()

	wr := float64(width) / float64(imgw)
	wh := float64(height) / float64(imgh)
	scalingRatio := math.Min(wr, wh)
	// center
	xoffset := 0.0
	yoffset := 0.0
	if scalingRatio*float64(imgw) < float64(width) {
		xoffset = (float64(width) - scalingRatio*float64(imgw)) / 2
		res = append(res, &wordclouds.Box{
			float64(height),
			0.0,
			xoffset,
			0,
		})
		res = append(res, &wordclouds.Box{
			float64(height),
			float64(width) - xoffset,
			float64(width),
			0,
		})
	}

	if scalingRatio*float64(imgh) < float64(height) {
		yoffset = (float64(height) - scalingRatio*float64(imgh)) / 2
		res = append(res, &wordclouds.Box{
			yoffset,
			0.0,
			float64(width),
			0,
		})
		res = append(res, &wordclouds.Box{
			float64(height),
			0.0,
			float64(width),
			float64(height) - yoffset,
		})
	}
	step := 3
	bounds := img.Bounds()
	for i := bounds.Min.X; i < bounds.Max.X; i = i + step {
		for j := bounds.Min.Y; j < bounds.Max.Y; j = j + step {
			r, g, b, a := img.At(i, j).RGBA()
			er, eg, eb, ea := exclude.RGBA()

			if r == er && g == eg && b == eb && a == ea {
				b := &wordclouds.Box{
					math.Min(float64(j+step)*scalingRatio+yoffset, float64(height)),
					float64(i)*scalingRatio + xoffset,
					math.Min(float64(i+step)*scalingRatio+xoffset, float64(width)),
					float64(j)*scalingRatio + yoffset,
				}
				res = append(res, b)
			}
		}
	}

	return res
}
