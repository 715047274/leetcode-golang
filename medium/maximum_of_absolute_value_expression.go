package medium

// Given two arrays of integers with equal lengths, return the
// maximum value of:
// |arr1[i] - arr1[j]| + |arr2[i] - arr2[j]| + |i - j|
// where the maximum is taken over all 0 <= i, j < arr1.length.
//
// Example 1:
// Input: arr1 = [1,2,3,4], arr2 = [-1,4,5,6]
// Output: 13
//
// Example 2:
// Input: arr1 = [1,-2,-5,0,10], arr2 = [0,-2,-1,-7,-4]
// Output: 20
//
// Constraints:
// 2 <= arr1.length == arr2.length <= 40000
// -10^6 <= arr1[i], arr2[i] <= 10^6

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	size := len(arr1)
	max := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			cur := abs(arr1[i]-arr1[j]) + abs(arr2[i]-arr2[j]) + abs(i-j)
			if cur > max {
				max = cur
			}
		}
	}
	return max
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
