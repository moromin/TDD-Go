package sort_test

import (
	"go-tdd/utils/sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Init
	elements := []int{4, 1, 6, 5, 2, 3, 8, 7, 9}

	// Execution
	sort.BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Errorf("first element should be 9")
	}
	if elements[len(elements)-1] != 1 {
		t.Errorf("last element should be 1")
	}
}
