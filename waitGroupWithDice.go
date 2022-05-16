package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go playDiceFor(&waitGroup, 3, "India")
	go playDiceFor(&waitGroup, 3, "Pakistan")
	waitGroup.Wait()

}

func playDiceFor(wg *sync.WaitGroup, breakValue int, playerName string) {
	//Generate a random number x where x is in range 5<=x<=20
	rangeLower := 1
	rangeUpper := 5
	counter := 1
	for {
		rand.Seed(time.Now().UnixNano())
		randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
		//
		if randomNum == breakValue {
			fmt.Println("**player:", playerName, " found the match in ", counter, " attempt")
			break
		} else {
			fmt.Println(playerName, " going to sleep for ", randomNum)

			time.Sleep(time.Second * time.Duration(randomNum))
		}
		counter++
	}
	wg.Done()

}

// run with command
// go run waitGroup2.go
