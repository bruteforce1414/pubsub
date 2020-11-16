package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewItem(t *testing.T) {

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