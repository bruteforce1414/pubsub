package main

import "github.com/bruteforce1414/pubsub/observer"

func main() {

	shirtItem := newItem("Nike")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()
}
