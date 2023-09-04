func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    var ans []int
    for idx, num := range nums {
        f := target - num
        v, ok := m[f]
        if ok {
            ans = append(ans, v)
            ans = append(ans, idx)
            return ans
        }
        m[num] = idx
    }

    return ans
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if val, found := m[target-num]; found {
			return []int{val, idx}
		}

		m[num] = idx
	}
	return nil
}
