package main

import "github.com/bruteforce1414/pubsub/observer"

func main() {

	shirtItem := observer.NewItem("Nike shirt")

	observerFirst := &observer.Customer{Id: "abc@gmail.com"}
	observerSecond := &observer.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
}
