package bubblesort

import (
	"github.com/mastodilu/go-sort/comparable"
)

// items contains the pointer to the slice of items to sort
var items *[]comparable.Comparable

// Sort sorts the given slice of Comparable using bubblesort algorithms
func Sort(slice *[]comparable.Comparable) {
	items = slice

	// len < 2: already sorted
	if len(*items) < 2 {
		return
	}

	// len >= 2
	bubblesort()
}

func bubblesort() {
	// index of the first item to the right that must be sorted
	//                          last
	//                           ⬇
	// X X X X X X X X X X  ...  2 8 9
	last := len(*items) - 1

	for last > 0 {
		swapped := false
		for i := 0; i < last; i++ {
			if (*items)[i+1].Less((*items)[i]) {
				swap(i, i+1)
				swapped = true
			}
		}
		last--
		
		// can stop when the the slice is sorted: when after a complete visit no swap is done
		if !swapped {
			return
		}
	}
}

// swap swaps the content of the items at the two indexes
func swap(i, j int) {
	temp := (*items)[i]
	(*items)[i] = (*items)[j]
	(*items)[j] = temp

}
