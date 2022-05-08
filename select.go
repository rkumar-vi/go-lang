package main

import (
	"fmt"
	"time"
)

func main() {

	channel_1 := make(chan string)

	go portal_1(channel_1)

	select {
	case message_from_channel_1 := <-channel_1:
		fmt.Println(message_from_channel_1)
	case exit := <-channel_1:
		fmt.Println(exit)
	}

}

func portal_1(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "welcome to channel_1"
}
