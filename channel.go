package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	channel := make(chan string)
	go playDice(channel)

	//for {
	channelOutput := <-channel
	fmt.Println(channelOutput)
	//}
}

func playDice(channel chan string) {
	var result string = " " //:= str + strconv.Itoa(n)
	for {
		randomInteger := rand.Intn(6) - 1
		result = result + strconv.Itoa(randomInteger)
		if randomInteger == 3 {
			fmt.Println(result)
			close(channel)
			return
		} else {
			fmt.Println(result)
		}
		//channel <- randomInteger
	}

}
