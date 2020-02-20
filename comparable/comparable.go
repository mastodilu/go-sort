package comparable

type Comparable interface {
	Less(Comparable) bool
	Equals(Comparable) bool
}
