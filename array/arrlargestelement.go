// Package array demonstrates solutions to various array problems
package array

// LargestElement finds largest element in a passed in integer array
//
// problem:
//
// Given an array A[] of size n. The task is to find the largest element in it.
//
// Expected Time Complexity: O(N)
//
// Expected Auxiliary Space: O(1)
//
// constraints:
//
// 1 <= n <= 1000
//
// 0 <= A[i] <= 1000
//
// Array may contain duplicate elements.
//
// author(s) [Dan Cochran](https://github.com/dancochran)
func LargestElement(arr []int, size int) int {

	largest := arr[0]

	for i := 1; i < size; i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}

	return largest
}
