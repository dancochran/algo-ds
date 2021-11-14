// arrlargestelement_test.go
// description: Test for finding largest element in array
// author(s) [Dan Cochran](https://github.com/dancochran)
// see arrlargestelement.go

package array

import "testing"

func TestLargestElementArrayOne(t *testing.T) {
	got := LargestElement([]int{66}, 1)
	want := 66

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLargestElementArraySeven(t *testing.T) {
	got := LargestElement([]int{39, 24, 6, 9, 91, 345, 84}, 7)
	want := 345

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
