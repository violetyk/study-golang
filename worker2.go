// +build ignore

package main

import (
	"fmt"
	"log"
	"sync"
)

// func worker(msg string) (<-chan string, <-chan bool) {
func worker(msg string) <-chan string {
	var wg sync.WaitGroup
	receiver := make(chan string)
	// fin := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(i int) {
				msg := fmt.Sprintf("%d %s done", i, msg)
				receiver <- msg
				wg.Done()
			}(i)
		}
		wg.Wait()
		// fin <- false // 終了を伝える
		// チャンネルを閉じるとfalseとからメッセージが返ることを利用して上と同じことをする
		close(receiver)
	}()
	// return receiver, fin
	return receiver
}

func main() {
	// 用途が違う読み取りチャンネルを受信してselectで処理を分ける
	// receiver, fin := worker("job")
	// for {
	// select {
	// case receive := <-receiver:
	// log.Println(receive)
	// case <-fin: // 終了したら終わる
	// return
	// }
	// }

	// チャンネルのclose()を利用する
	receiver := worker("job")
	for {
		// closeするとokにfalseが返る
		receive, ok := <-receiver
		if !ok {
			log.Println("closed")
			return
		}
		log.Println(receive)
	}
}
