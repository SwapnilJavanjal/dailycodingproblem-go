package day102

// ContiguousSumBrute uses brute-force to return the
// contiguous subset that sums to K.
// Runtime is O(N^2) and O(1) space.
func ContiguousSumBrute(nums []int, k int) []int {
	var result []int
	for i := range nums {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				return nums[i : j+1]
			}
		}
	}
	return result
}

// ContiguousSumNonNegative uses a window to return the
// contiguous subset that sums to K.
// Can only tolerate non-negative values, panics otherwise.
// Runtime is O(N) and O(1) space.
func ContiguousSumNonNegative(nums []int, k int) []int {
	var result []int
	var sum, begin int
	for end := 0; end < len(nums); end++ {
		if nums[end] < 0 {
			panic("negative values aren't permitted")
		}
		sum += nums[end]
		if sum == k {
			result = nums[begin : end+1]
			break
		} else if sum > k {
			sum -= nums[begin]
			sum -= nums[end]
			begin++
			end--
		}
	}
	return result
}

// ContiguousSumNegatives returns the
// contiguous subset that sums to K.
// This implementation tolerates negative values.
// Runtime is O(N) and O(N) space.
func ContiguousSumNegatives(nums []int, k int) []int {
	var result []int
	var sum int
	m := make(map[int]int)
	for i := range nums {
		sum += nums[i]
		if sum-k == 0 {
			result = nums[0 : i+1]
			break
		}
		if start, found := m[sum-k]; found {
			result = nums[start+1 : i+1]
			break
		}
		m[sum] = i
	}
	return result
}
