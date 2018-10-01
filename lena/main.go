package main

import (
	"fmt"
	"image"
	"os"

	"golang.org/x/image/bmp"
)

func saveImage(m image.Image, filename string) {
	writeF, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer writeF.Close()

	err = bmp.Encode(writeF, m)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("It's Lena!!")

	f, err := os.Open("./lena.bmp")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	i, err := bmp.Decode(f)
	if err != nil {
		panic(err)
	}

	leftRightImg := image.NewGray(image.Rect(0, 0, 512, 512))
	upDownImg := image.NewGray(image.Rect(0, 0, 512, 512))
	mirrorImg := image.NewGray(image.Rect(0, 0, 512, 512))

	for n := 0; n < i.Bounds().Max.X; n++ {
		for m := 0; m < i.Bounds().Max.Y; m++ {
			c := i.At(n, m)
			leftRightImg.Set(i.Bounds().Max.X-n, m, c)
			upDownImg.Set(n, i.Bounds().Max.Y-m, c)
			mirrorImg.Set(i.Bounds().Max.X-m, i.Bounds().Max.Y-n, c)
		}
	}

	saveImage(leftRightImg, "leftright.bmp")
	saveImage(upDownImg, "updown.bmp")
	saveImage(mirrorImg, "mirror.bmp")
}
