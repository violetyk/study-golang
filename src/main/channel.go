package main

import (
	"log"
	"time"
)

func f(ch chan bool) {
	log.Println("waiting")
	time.Sleep(time.Second)
	ch <- true
}

func main() {
	ch := make(chan bool)
	go f(ch)
	log.Println(<-ch) // ここでデータが来るまでブロックする

}
