package comparableint

import (
	"github.com/mastodilu/go-sort/comparable"
)

// ComparableInt implements Comparable interface
type ComparableInt int64

// Equals returns true if the two items have the same value
func (ci ComparableInt) Equals(c comparable.Comparable) bool {
	ci2, ok := c.(ComparableInt)
	if !ok {
		panic("asfasdfadfadajhajldhf") // FIXME panic message
	}

	return ci == ci2
}

// Less returns true if the first items is less than the second item
func (ci ComparableInt) Less(c comparable.Comparable) bool {
	ci2, ok := c.(ComparableInt)
	if !ok {
		panic("asfasdfadfadajhajldhf") // FIXME panic message
	}

	return ci < ci2
}
