package main

import (
	"fmt"
	"log"
	"os"

	"image/jpeg"
	"image/png"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\t%s png_path jpg_path\n", os.Args[0])
		os.Exit(1)
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	reader, err := os.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	writer, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(writer, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
