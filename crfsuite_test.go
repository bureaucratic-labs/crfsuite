package crfsuite

import (
	"fmt"
	"testing"
)

func TestModelFromFile(t *testing.T) {
	model := NewModelFromFile("test_data/tokenization-model.crfsuite")
	labels := model.GetLabels()
	if labels.Length() != 3 {
		t.Fail()
	}
	ids := []int{
		labels.ToID("B"),
		labels.ToID("I"),
		labels.ToID("O"),
	}
	fmt.Println("%v", ids)
}
