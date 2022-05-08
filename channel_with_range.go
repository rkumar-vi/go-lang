package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	channel := make(chan string)
	go playDice(channel)

	for msg := range channel {
		fmt.Println(msg)
	}
}

func playDice(channel chan string) {
	var result string = " "
	for {
		randomInteger := rand.Intn(6) - 1
		result = strconv.Itoa(randomInteger)
		channel <- result
		if randomInteger == 3 {
			close(channel)
			return
		}

	}

}
