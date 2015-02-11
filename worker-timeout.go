package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func randomTime() time.Duration {
	return time.Duration(rand.Intn(1e3)) * time.Millisecond
}

func worker(msg string) <-chan string {
	receiver := make(chan string)
	for i := 0; i < 300; i++ {
		go func(i int) {
			time.Sleep(randomTime())
			msg := fmt.Sprintf("%d %s done", i, msg)
			receiver <- msg
		}(i)
	}
	return receiver
}

func main() {
	receiver := worker("job")
	for {
		select {
		case receive := <-receiver:
			log.Println(receive)
			// func After(d Duration) <-chan Time
			//  一定時間経過したらメッセージを送るチャンネルを返す
		case <-time.After(time.Second): // 1秒以上立ったら受け取るメッセージ（＝重い処理）
			log.Println("timeout")
			return // 受信時に終われば timeout 処理になる。
		}
	}
}
