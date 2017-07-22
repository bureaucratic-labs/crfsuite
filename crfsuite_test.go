package crfsuite

import (
	"fmt"
	"testing"
)

func TestModelFromFile(t *testing.T) {
	model := NewModelFromFile("test_data/tokenization-model.crfsuite")
	labels := model.Labels
	if labels.Length() != 3 {
		t.Fail()
	}
	ids := []int{
		labels.ToID("B"),
		labels.ToID("I"),
		labels.ToID("O"),
	}
	fmt.Println("%v", ids)
	instance := NewInstance()
	item := NewItem()
	attribute := NewAttribute(1, 1.0)
	fmt.Println("%v", attribute)
	item.AddAttribute(attribute)
	instance.AddItem(item, labels.ToID("B"))
}
