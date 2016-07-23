package data_structures

import "testing"

func TestContainsElementsAddedToLinkedList(t *testing.T) {
	
	items := [5]int{11,22,33,44,55}
	
	for _,item:= range items{
		Add(item)
	}

	for _,item:= range items{
		if Search(item) == nil {
			t.Errorf("Should have contained a node with value %d", item)
		}
	}
	
	/*Add(1)
	Add(2)
	Add(3)

	for i := 1; i <= 3; i++ {
		if Search(i) == nil {
			t.Errorf("Should have contained a node with value %d", i)
		}
	}*/
}
