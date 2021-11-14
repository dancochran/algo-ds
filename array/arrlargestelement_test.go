package array

import "testing"

// TestLargestElementArrayOne tests finding the largest element in array of size one
// author(s) [Dan Cochran](https://github.com/dancochran)
func TestLargestElementArrayOne(t *testing.T) {
	got := LargestElement([]int{66}, 1)
	want := 66

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// TestLargestElementArraySeven tests finding the largest element in array of size seven
// author(s) [Dan Cochran](https://github.com/dancochran)
func TestLargestElementArraySeven(t *testing.T) {
	got := LargestElement([]int{39, 24, 6, 9, 91, 345, 84}, 7)
	want := 345

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
