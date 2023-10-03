func canJump(nums []int) bool {
    if len(nums) == 1 {
		return true
	}

	idx := 0

	for true {
		var jump = nums[idx]
		if jump == 0 { return false }
		if idx + jump >= len(nums) - 1 { return true }

		var k = 1
		var m = math.MinInt32
		var next = 0
		for k <= jump {
			if nums[idx + k] + k + idx > m {
				next = k + idx
				m = nums[k + idx] + k + idx
			}
			k++
		}

		idx = next
	}
}