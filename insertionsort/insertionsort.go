package insertionsort

import (
	"log"

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
	if left > right {
		log.Fatalf("error: left > right (%v > %v)\n", left, right)
	}

	temp := (*items)[right]

	for i := right; i > left; i-- {
		(*items)[i] = (*items)[i-1]
	}
	(*items)[left] = temp
}
