package two_sum

import "slices"

func twoSum(nums []int, target int) []int {
	var (
		result       []int
		firstNumber  int
		secondNumber int
		found        bool
	)

	numsSorted := make([]int, len(nums))
	copy(numsSorted, nums)

	// O(n*logn)
	slices.Sort(numsSorted)

	// O(n*logn)
	for i := 0; i < len(numsSorted); i++ {
		currentTarget := target - numsSorted[i]
		low, high := 0, len(numsSorted)-1

		for low <= high {
			mid := low + (high-low)/2

			if currentTarget < numsSorted[mid] {
				high = mid - 1
			} else if currentTarget > numsSorted[mid] {
				low = mid + 1
			} else {
				if i != mid {
					firstNumber = numsSorted[i]
					secondNumber = numsSorted[mid]
					found = true
				}
				break
			}
		}

		if found {
			break
		}
	}

	// O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i] == firstNumber || nums[i] == secondNumber {
			result = append(result, i)
		}
	}

	return result
}
