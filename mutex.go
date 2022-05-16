package main

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}

func main() {

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup, &mutex)
	}
	waitGroup.Wait()
	fmt.Println("sum of 1 thousand times = ", x)

}
