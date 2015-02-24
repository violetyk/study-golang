// +build ignore

package main

import (
	"bytes"
	"image/jpeg"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"

	"github.com/disintegration/imaging"
	"github.com/k0kubun/pp"
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
	// detect format
	if reflect.DeepEqual(data[0:2], []byte{0xff, 0xd8}) {
		pp.Println("Format: jpeg")
	}

	config, err := jpeg.DecodeConfig(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	orig_width := config.Width
	orig_height := config.Height

	img, err := jpeg.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	resized := imaging.Resize(img, orig_width/2, orig_height/2, imaging.Lanczos)

	dist_path := "/Users/yuhei/Desktop/takao2_resized_by_imaging.jpg"
	file2, err := os.Create(dist_path)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	imaging.Encode(file2, resized, imaging.JPEG)
}
