func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
	len := len(nums);
	for i := 0; i < len; i++ {
		_, prs := m[nums[i]]
		if prs {
			return true
		}
		m[nums[i]] = 1;
	}

	return false
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	xm := make(map[int]struct{}) // empty struct

	for _, v := range nums {
		if _, ok := xm[v]; ok {
			return true
		}

		xm[v] = struct{}{}
	}

	return false
}