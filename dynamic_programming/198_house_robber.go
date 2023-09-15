func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    m := make([]int, len(nums) + 1)
    m[0] = 0
    m[1] = nums[0]
    for i := 2; i < len(nums) + 1; i++ {
        if m[i - 1] > m[i - 2] + nums[i - 1] {
            m[i] = m[i - 1]
        }else {
            m[i] = m[i - 2] + nums[i - 1]
        }
    }

    return m[len(nums)]
}