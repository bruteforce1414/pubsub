package main

import (
	"fmt"
	"github.com/bruteforce1414/pubsub/observer"
	"github.com/bruteforce1414/pubsub/customer"
)

func main() {

	shirtItem := observer.NewItem("Nike")

	observerFirst := &customer.Customer{Id: "abc@gmail.com"}
	observerSecond := &customer.Customer{Id: "xyz@gmail.com"}

	fmt.Println("shirtItem", shirtItem)
	fmt.Println("observerFirst", observerFirst)
	fmt.Println("observerSecond", observerSecond)


	//shirtItem.Register(observerFirst)
	//shirtItem.Register(observerSecond)
	//shirtItem.UpdateAvailability()
}
