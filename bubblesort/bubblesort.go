package bubblesort

import (
	"fmt"

	"github.com/mastodilu/go-sort/comparable"
)

// items contains the pointer to the slice of items to sort
var items *[]comparable.Comparable

// Sort sorts the given slice of Comparable using bubblesort algorithms
func Sort(slice *[]comparable.Comparable) error {
	items = slice
	if len(*items) < 2 {
		return nil
	}
	return bubblesort()
}

func bubblesort() error {
	// for i := 0; i < len(*items-1); i++ {

	// }
	return fmt.Errorf("method not implemented yet")
}

// swap swaps the content of the items at the two indexes
func swap(i, j int) {
	temp := (*items)[i]
	(*items)[i] = (*items)[j]
	(*items)[j] = temp

}
