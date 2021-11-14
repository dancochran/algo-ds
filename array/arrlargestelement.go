// arrlargestelement.go
// description: Finds largest element in an integer array
// problem::
// Given an array A[] of size n. The task is to find the largest element in it.
//
// Expected Time Complexity: O(N)
// Expected Auxiliary Space: O(1)
//
// constraints:
// 1 <= n <= 1000
// 0 <= A[i] <= 1000
// Array may contain duplicate elements.
//
// author(s) [Dan Cochran](https://github.com/dancochran)
// see arithmeticmean_test.go

// Package array demonstrates solutions to varuous array problems
package array

func LargestElement(arr []int, size int) int {

	largest := arr[0]

	for i := 1; i < size; i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}

	return largest
}
