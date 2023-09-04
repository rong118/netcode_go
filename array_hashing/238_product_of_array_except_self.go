import (
    "fmt"
)

func productExceptSelf(nums []int) []int {
    var l = len(nums)
    left := make([]int, l + 1)
    left[0] = 1
    for i := 0; i < l; i++ {
        left[i + 1] = left[i] * nums[i]
    }

    right := make([]int, l + 1)
    right[l] = 1;
    for i := l - 1; i >= 0; i-- {
        right[i] = right[i + 1] * nums[i]
    }

    ans := make([]int, l)
    for i := 0; i < l; i++ {
        ans[i] = left[i] * right[i + 1];
    }

    return ans
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
