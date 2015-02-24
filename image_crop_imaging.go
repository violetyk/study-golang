// +build ignore

package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/disintegration/imaging"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	orig_path := "/var/tmp/graid/raw.githubusercontent.com/violetyk/graid/master/test_data/takao2.jpg/default"

	file, err := os.Open(orig_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	// w, h := 400, 300
	// resized := imaging.CropCenter(img, w, h)
	x0, y0, x1, y1 := 200, 300, 100, 50 // x:100, y:50の１からx:200, y:300のところまでなので、width=100, height=250
	resized := imaging.Crop(img, image.Rect(x0, y0, x1, y1))

	dist_path := "/Users/yuhei/Desktop/takao2_cropped_by_imaging.jpg"
	file2, err := os.Create(dist_path)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	imaging.Encode(file2, resized, imaging.JPEG)

}
