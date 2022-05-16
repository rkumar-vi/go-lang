package main

import (
	"fmt"
)

func main() {

	channel_1 := make(chan string)
	channel_2 := make(chan string)

	go portal_1(channel_1)
	go portal_2(channel_2)

	for {
		select {
		case message_from_channel_1 := <-channel_1:
			fmt.Println(message_from_channel_1)
			//go portal_2(channel_2)
		case message_from_channel_2 := <-channel_2:
			fmt.Println(message_from_channel_2)
			//go portal_1(channel_1)
		}
	}
}

func portal_1(channel chan string) {
	//time.Sleep(1 * time.Second)
	channel <- "welcome to channel_1"
}
func portal_2(channel chan string) {
	//time.Sleep(3 * time.Second)
	channel <- "welcome to channel_2"
}
