// +build ignore

package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	orig_path := "/var/tmp/graid/raw.githubusercontent.com/violetyk/graid/master/test_data/takao1.jpg/default"
	// dist_path := "/Users/yuhei/Desktop/resized.jpg"
	dist_path := "/Users/yuhei/Desktop/thumbnail.jpg"

	file, err := os.Open(orig_path)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	var width uint = 400
	var height uint = 0
	m := resize.Resize(width, height, img, resize.Lanczos3)
	// m := resize.Thumbnail(width, height, img, resize.Lanczos3)

	out, err := os.Create(dist_path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

}
