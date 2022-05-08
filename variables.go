/*
## May-2022
This file will be used to show different ways to declare and initialise variables in GOLANG
command to run this programme is "go run variables.go"
*/

package main

import "fmt"

func main() {

	var temperature_in_celcius int // declare variable
	temperature_in_celcius = 30    // initialise variable

	fmt.Println("Temperature in celcius is ", temperature_in_celcius)

	temperature_in_fahrenheight := 90 // declaration and initialisation together. This is called Short Hand declaration
	fmt.Println("Temperature in Fahrenhieght is ", temperature_in_fahrenheight)

	var (
		length_of_room_in_feet = 8
		width_of_room_in_feet  = 10
	)
	fmt.Println("\nArea of room is ", length_of_room_in_feet*width_of_room_in_feet, " sq feet")
}
