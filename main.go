package main

import (
	"github.com/bruteforce1414/pubsub/observer"
	"github.com/bruteforce1414/pubsub/customer"
)

func main() {

	shirtItem := observer.NewItem("Nike")

	observerFirst := &customer.Customer{id: "abc@gmail.com"}
	observerSecond := &customer.Customer{id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
}
