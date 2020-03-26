package insertionsort

import (
	"testing"

	"github.com/mastodilu/go-sort/comparable"
	"github.com/mastodilu/go-sort/comparableint"
)

func TestSort(t *testing.T) {
	temp := []comparable.Comparable{
		comparableint.ComparableInt(8),
		comparableint.ComparableInt(1),
		comparableint.ComparableInt(2),
		comparableint.ComparableInt(0),
		comparableint.ComparableInt(5),
		comparableint.ComparableInt(6),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(9),
		comparableint.ComparableInt(7),
	}
	expected := []comparable.Comparable{
		comparableint.ComparableInt(0),
		comparableint.ComparableInt(1),
		comparableint.ComparableInt(2),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(5),
		comparableint.ComparableInt(6),
		comparableint.ComparableInt(7),
		comparableint.ComparableInt(8),
		comparableint.ComparableInt(9),
	}

	Sort(&temp)

	for i, c := range expected {
		if !temp[i].Equals(c) {
			t.Errorf("index %d, expected %d, got %d", i, c, temp[i])
		}
	}

}
