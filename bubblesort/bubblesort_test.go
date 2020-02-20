package bubblesort

import (
	"testing"

	"github.com/mastodilu/go-sort/comparable"
	"github.com/mastodilu/go-sort/comparableint"
)

func TestSwap(t *testing.T) {
	temp := []comparable.Comparable{
		comparableint.ComparableInt(1), // <-- 1
		comparableint.ComparableInt(2),
		comparableint.ComparableInt(3), // <-- 2
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(5),
	}

	items = &temp // sets the slice used by swap(..)
	i, j := 0, 2

	expected := []comparable.Comparable{
		comparableint.ComparableInt(3), // <-- 2
		comparableint.ComparableInt(2),
		comparableint.ComparableInt(1), // <-- 1
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(5),
	}

	swap(i, j)

	for i, item := range expected {
		if item != temp[i] {
			t.Errorf("index %d, expected %d, got %d", i, item, temp[i])
		}
	}

}

func TestBubblesort(t *testing.T) {

	items = &([]comparable.Comparable{
		comparableint.ComparableInt(0),
		comparableint.ComparableInt(5),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(2),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(5),
		comparableint.ComparableInt(5),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(1),
	})

	expected := []comparable.Comparable{
		comparableint.ComparableInt(0),
		comparableint.ComparableInt(1),
		comparableint.ComparableInt(2),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(5),
		comparableint.ComparableInt(5),
		comparableint.ComparableInt(5),
	}

	bubblesort()

	for i, item := range *items {
		if !item.Equals(expected[i]) {
			t.Errorf("index %d: expected %v, got %v", i, expected[i], item)
		}
	}

	// test n2 ------------------------

	items = &[]comparable.Comparable{
		comparableint.ComparableInt(5),
	}

	expected = []comparable.Comparable{
		comparableint.ComparableInt(5),
	}

	bubblesort()

	for i, item := range *items {
		if !item.Equals(expected[i]) {
			t.Errorf("index %d: expected %v, got %v", i, expected[i], item)
		}
	}

	// test n3 ------------------------

	items = &[]comparable.Comparable{}

	expected = []comparable.Comparable{}

	bubblesort()

	for i, item := range *items {
		if !item.Equals(expected[i]) {
			t.Errorf("index %d: expected %v, got %v", i, expected[i], item)
		}
	}
}
