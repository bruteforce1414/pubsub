package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewItem(t *testing.T) {

	type itemForCompare struct {
		observerList []observer
		name         string
		inStock      bool
	}

	itemForCompare1:=&itemForCompare{
		observerList: nil,
		name:         "Shirt Adidas",
		inStock:      false,
	}

	testItem1 := NewItem("Shirt Adidas")
	testItem2 := NewItem("Shirt Reebok")
	assert.Equal(t, testItem1.name, itemForCompare1.name)
	assert.NotEqual(t, testItem2.name, itemForCompare1.name)
}

func Test_Register(t *testing.T){
	testItem1 := NewItem("Nike shirt")
	observerTest1 := &Customer{Id: "abc@gmail.com"}
	observerTest2 := &Customer{Id: "xyz@gmail.com"}
	testItem1.Register(observerTest1)
	testItem1.Register(observerTest2)

	assert.Equal(t,testItem1.observerList[0].getID(), "abc@gmail.com");
	assert.Equal(t,testItem1.observerList[1].getID(), "xyz@gmail.com");
	assert.Equal(t, len(testItem1.observerList), 2);

}

func Test_Deregister(t *testing.T){
	testItem1 := NewItem("Nike shirt")
	observerTest1 := &Customer{Id: "abc@gmail.com"}
	observerTest2 := &Customer{Id: "xyz@gmail.com"}
	testItem1.Register(observerTest1)
	testItem1.Register(observerTest2)
	testItem1.Deregister(observerTest2)

	assert.Equal(t,testItem1.observerList[0].getID(), "abc@gmail.com");
	assert.Equal(t, len(testItem1.observerList), 1);
}