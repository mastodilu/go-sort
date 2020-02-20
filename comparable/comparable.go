package comparable

// Comparable represents a type that can be compared
type Comparable interface {
	Less(Comparable) bool
	Equals(Comparable) bool
}
