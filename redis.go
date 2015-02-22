// +build ignore

package main

import (
	"bytes"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/k0kubun/pp"
)

func main() {
	host := "localhost"
	port := "6379"
	c, err := redis.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// set string
	c.Do("SET", "key", "value1234567890")

	// get string
	s, err := redis.String(c.Do("GET", "key"))
	if err != nil {
		panic(err)
	}
	pp.Print(s)

	c.Flush()

	// set image
	// read from file
	image_path := "/var/tmp/graid/raw.githubusercontent.com/violetyk/graid/master/test_data/takao2.jpg/default"
	file, err := os.Open(image_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// image, err := jpeg.Decode(file)
	// if err != nil {
	// log.Fatal(err)
	// }
	// file.Close()

	// set image to redis
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	c.Do("SET", "takao2.jpg", bs)

	// get image
	// image from byte slice
	data, err := redis.Bytes(c.Do("GET", "takao2.jpg"))
	if err != nil {
		panic(err)
	}
	image2, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	dist_path := "/Users/yuhei/Desktop/read_from_redis.jpg"
	out, err := os.Create(dist_path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// jpeg.Encode(out, image2, nil)
	// set quality
	jpeg.Encode(out, image2, &jpeg.Options{Quality: 100})

}
