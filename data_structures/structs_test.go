package data_structures

import "testing"

func TestContainsElementsAddedToLinkedList(t *testing.T) {
	Add(1)
	Add(2)
	Add(3)

	for i := 1; i <= 3; i++ {
		if Search(i) == nil {
			t.Errorf("Should have contained a node with value %d", i)
		}
	}
}
