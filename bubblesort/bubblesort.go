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
	return fmt.Errorf("bubblesort.Sort() method not implemented yet")
}
