package utils

import (
	"testing"

	"github.com/mastodilu/go-sort/comparable"
	"github.com/mastodilu/go-sort/comparableint"
)

func TestCheckSorted(t *testing.T) {
	slice := &[]comparable.Comparable{
		comparableint.ComparableInt(0),
		comparableint.ComparableInt(2),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(18),
	}

	err := CheckSorted(slice)
	if err != nil {
		t.Errorf(err.Error())
	}

	slice = &[]comparable.Comparable{
		comparableint.ComparableInt(0),
		comparableint.ComparableInt(3),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(4),
		comparableint.ComparableInt(18),
		comparableint.ComparableInt(2),
	}

	err = CheckSorted(slice)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
