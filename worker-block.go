// +build ignore

// チャネルバッファはメッセージキューみたいなもの
// バッファがいっぱいだったらメッセージを送信しないことを利用して
// ワーカーの同時期同数を制限する

package main

import (
	"fmt"
	"log"
	"runtime"
)

func worker(msg string) <-chan string {
	limit := make(chan int, 5)
	receiver := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			// 必ずしも5にはならない
			log.Println(runtime.NumGoroutine())
			limit <- 1
			go func(i int) {
				msg := fmt.Sprintf("%d %s done", i, msg)
				receiver <- msg
				<-limit
			}(i)
		}
	}()
	return receiver
}

func main() {
	receiver := worker("job")
	for i := 0; i < 100; i++ {
		log.Println(<-receiver)
	}
}
