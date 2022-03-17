package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const consumerCount int = 3

var messages = []string{
	"The world itself's",
	"just one big hoax.1",
	"just one big hoax.2",
	"just one big hoax.3",
	"just one big hoax.4",
	"just one big hoax.5",
}


func produce(jobs chan string,ticker *time.Ticker,wg *sync.WaitGroup) {
	defer wg.Done()
	for{
		select {
			case <- ticker.C:
				println("========================================================")
				for _, msg := range messages {
					jobs <- msg
				}
		}
	}


}

func consume(worker int, jobs chan string, done chan bool,wg *sync.WaitGroup) {

	for true {
		select {
		case job := <-jobs:
			if worker == 2 {
				time.Sleep(1000)
			}

			fmt.Printf("Message %v is consumed by worker %v.\n", job, worker)
		case <- done:
			println("DOneeeeeeeeeeeeeeeeeee")
			wg.Done()
			os.Exit(1)
			return
		}

	}

}

func main() {
	jobs := make(chan string,50)
	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)

	var wg sync.WaitGroup
	go produce(jobs,ticker,&wg)
	for i := 1; i <= consumerCount; i++ {
		wg.Add(1)
		go consume(i, jobs, done,&wg)
	}
	time.Sleep(10 * time.Second)
	done<-true
	wg.Wait()
	print("Done")

}