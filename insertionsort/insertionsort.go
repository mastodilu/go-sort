package insertionsort

import (
	"github.com/mastodilu/go-sort/comparable"
)

// items contains the pointer to the slice of items to sort
var items *[]comparable.Comparable

// Sort calls the insertion sort algorithm
func Sort(slice *[]comparable.Comparable) {
	items = slice
	if items != nil && len(*items) > 1 {
		insertionSort()
	}
}

func insertionSort() {
	for i := 1; i < len(*items); i++ {
		insert(i)
	}
}

func insert(index int) {
	for i := 0; i < len((*items)[:index]); i++ {
		if (*items)[index].Less((*items)[i]) {
			insertAndShift(i, index)
			return
		}
	}
}

// insertAndShift inserts the item at index 'right'
// before the item at index 'left'
func insertAndShift(left, right int) {
	temp := (*items)[right]

	// remove the item we want to insert from its original position
	*items = append(
		(*items)[0:right],
		(*items)[right+1:]...,
	)

	if left == 0 {
		// insert before the first item
		*items = append(
			[]comparable.Comparable{temp},
			*items...,
		)
	} else {
		// insert in the middle
		*items = append(
			(*items)[:left], // the LEFT part of the slice
			append(
				[]comparable.Comparable{temp}, // the ITEM to insert
				(*items)[left:]...,            // the RIGHT part of the array
			)...,
		)
	}
}
