// +build ignore

package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/disintegration/gift"
	"github.com/disintegration/imaging"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	orig_path := "/Users/kagaya/Desktop/test.jpg"

	file, err := os.Open(orig_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	src, err := jpeg.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	g := gift.New()
	g.Empty()
	g.Add(gift.Resize(200, 0, gift.LanczosResampling))

	dst := image.NewRGBA(g.Bounds(src.Bounds()))

	g.Draw(dst, src)

	dist_path := "/Users/kagaya/Desktop/takao2_resized_by_gift.jpg"
	file2, err := os.Create(dist_path)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	imaging.Encode(file2, dst, imaging.JPEG)
}
