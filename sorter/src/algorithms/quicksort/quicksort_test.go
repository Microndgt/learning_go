package quicksort

import "testing"

func TestQuickSort(t *testing.T){
	values := []int{4, 6, 9, 1, 5, 4, 3, 2, 1}
	sorted_values := QuickSort(values)
	if sorted_values[0] != 1 || sorted_values[1] != 1 || sorted_values[2] != 2 || sorted_values[3] != 3 || sorted_values[4] != 4{
		t.Error("QuickSort() failed. Got", sorted_values, "Expected 1 1 2 3 4")
	}
}

func TestQuickSort2(t *testing.T){
	values := []int{4, 6, 9, 1, 5, 4, 3, 2, 1}
	QuickSort2(values)
	if values[0] != 1 || values[1] != 1 || values[2] != 2 || values[3] != 3 || values[4] != 4{
		t.Error("QuickSort() failed. Got", values, "Expected 1 1 2 3 4")
	}
}