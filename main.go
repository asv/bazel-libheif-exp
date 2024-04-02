package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	//_ "bazel-libheif-exp/heif"
	_ "github.com/strukturag/libheif/go/heif"
)

func main() {
	var (
		outFilePath = flag.String("out", "", "")
		inFilePath  = flag.String("in", "", "")
	)
	flag.Parse()

	f, err := os.Open(*inFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, imgFormat, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	fo, err := os.OpenFile(*outFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	// Encode the image to JPEG.
	if err := jpeg.Encode(fo, img, &jpeg.Options{Quality: 95}); err != nil {
		panic(fmt.Errorf("encode %s image to jpeg: %w", imgFormat, err))
	}
}
