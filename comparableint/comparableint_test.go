package comparableint

import "testing"

func TestLess(t *testing.T) {
	first := ComparableInt(53)
	second := ComparableInt(34)

	got := second.Less(first)
	expected := true
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}

	third := ComparableInt(1)

	got = first.Less(third)
	expected = false
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestEquals(t *testing.T) {
	one := ComparableInt(1)
	two := ComparableInt(2)
	oneagain := ComparableInt(1)

	got := one.Equals(two)
	expected := false
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}

	got = one.Equals(oneagain)
	expected = true
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}
