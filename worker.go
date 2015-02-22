// +build ignore

package main

import (
	"fmt"
	"log"
)

// 読み取り専用のチャンネル<-chanを返す
func worker(msg string) <-chan string {
	receiver := make(chan string)
	// ワーカーを3つ起動
	for i := 0; i < 3; i++ {
		go func(i int) {
			msg := fmt.Sprintf("%d %s done", i, msg)
			receiver <- msg
		}(i)
	}
	return receiver
}

func main() {
	receiver := worker("job")

	// 先にワーカーを3つ起動したことが分かってるので
	// 3つ受け取るのを待ってる
	for i := 0; i < 3; i++ {
		// for i := 0; i < 4; i++ {
		log.Println(<-receiver)
	}
}
