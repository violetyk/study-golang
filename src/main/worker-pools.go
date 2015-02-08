package main

import (
	"log"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// ジョブが入ってきたらワーカーが実行、1秒待つ
	for j := range jobs {
		log.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 1.ワーカーを作って、ジョブはまだ与えない
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 2.ジョブを与える
	myjobs := 9
	for j := 1; j <= myjobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= myjobs; a++ {
		// 結果（今回は入らないので捨てるだけ）
		<-results
		// log.Println(<-results)
	}
}
