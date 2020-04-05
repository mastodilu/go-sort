package mergesort

import "github.com/mastodilu/go-sort/comparable"

// Sort runs mergesort if the collection has enough items
func Sort(c *[]comparable.Comparable) {
	if c != nil && len(*c) < 2 {
		return
	}

	*c = split(*c)
}

func split(items []comparable.Comparable) []comparable.Comparable {
	if len(items) < 2 {
		return items
	}

	// time.Sleep(time.Second * 2) // helps debugging

	center := int((len(items) - 1) / 2)

	// fmt.Printf("da %d a %d: %v\n", begin, end, (*collection)[begin:end+1])

	return merge(
		split(items[0:center+1]),
		split(items[center+1:]),
	)

}

func merge(left, right []comparable.Comparable) []comparable.Comparable {
	buffer := make([]comparable.Comparable, len(left)+len(right))
	buffIndex := 0

	for len(left) > 0 && len(right) > 0 {

		l, r := left[0], right[0]

		if l.Less(r) {
			buffer[buffIndex] = l
			left = left[1:]
		} else {
			buffer[buffIndex] = r
			right = right[1:]
		}
		buffIndex++
	}

	// append missing items
	if len(left) > 0 {
		for _, c := range left {
			buffer[buffIndex] = c
			buffIndex++
		}
	} else {
		for _, c := range right {
			buffer[buffIndex] = c
			buffIndex++
		}
	}

	return buffer
}
