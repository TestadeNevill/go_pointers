package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	//Reference to a variable use &
	changeUsingPointer(&myString)
	
	log.Println("after func call myString is set to", myString)
}
// Create a pointer type or location in memory *
func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}