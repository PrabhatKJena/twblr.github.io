package data_structures

import "testing"

func TestMapInts(t *testing.T) {
	square := func(arg1 int32) int32 {
		return arg1 * arg1
	}

	res := mapInts(square, []int32{1, 2, 3})
	expectedResult := []int32{1, 4, 9}

	if len(res) < 1 {
		t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)
	}

	for i := range res {
		if res[i] != expectedResult[i] {
			t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)
		}
	}

}

func TestFilterInts(t *testing.T) {
	greaterThanTwo := func(arg1 int32) bool {
		return arg1 > 2

	}

	res := filterInts(greaterThanTwo, []int32{1, 2, 3})
	expectedResult := []int32{3}

	if len(res) < 1 {
		t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)

	}

	for i := range res {
		if res[i] != expectedResult[i] {
			t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)
		}
	}
}

func TestConcatenation(t *testing.T) {
	superHeros := []string{"Batman", "Superman", "Flash"}

	newList := concatenate(superHeros, "Spiderman", "IronMan")

	if !equals(newList, []string{"Batman", "Superman", "Flash", "Spiderman", "IronMan"}) {
		t.Errorf("Concatenated list should have been equal")
	}
}

func TestSliceDoesNotCopy(t *testing.T) {
	marks := []int{23, 34, 56, 46, 100, 99, 24}

	partialReversedList := partialReverse(marks, 4, 6)
	expectedReversedList := []int{24, 99, 100}

	for idx, value := range expectedReversedList {
		if value != partialReversedList[idx] {
			t.Errorf("Expected '%d' at index '%d' but got '%d'", value, idx, partialReversedList)
		}
	}
}
