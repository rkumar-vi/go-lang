package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var masterGroup sync.WaitGroup
	masterGroup.Add(1)
	go worker(&masterGroup)
	masterGroup.Wait()
	fmt.Println(" After Wait() method. Ending main thread now")
}

func worker(masterCaller *sync.WaitGroup) {
	fmt.Println(" going to sleep for 10 seocnds")
	time.Sleep(time.Second * 10)
	fmt.Println(" woken up going to report  Task is done")
	masterCaller.Done()
}
