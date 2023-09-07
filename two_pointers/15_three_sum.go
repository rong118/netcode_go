func threeSum(nums []int) [][]int {
    // Sort the given array
	sort.Ints(nums)

    var ans [][]int
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        var l = i + 1
        var r = len(nums) - 1
        
        for l < r {
            var sum =  nums[i] + nums[l] + nums[r]
            if sum == 0 {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                l++
                for l < r && nums[l] == nums[l - 1] {
					l++
				}
            }else if sum < 0 {
                l++
            }else {
                r--
            }
        }
    }
    
    return ans;
}