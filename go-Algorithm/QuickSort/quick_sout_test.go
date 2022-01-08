package QuickSort

import "testing"

var tests = [][]int{
	{},
	{3, 2, 1},
	{1, 2, 5, 3, 11, 5, 6},
	{5, 4, 3, 2, 1},
}

func TestQuickSort(t *testing.T) {
	for _, test := range tests {
		sortedArray := QuickSort(test)
		for i := 0; i < len(sortedArray)-1; i++ {
			if test[i+1] < test[i] {
				t.Errorf("wrong!\n")
			}
		}
	}
}
