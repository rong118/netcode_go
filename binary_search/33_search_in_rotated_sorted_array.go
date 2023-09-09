func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left + 1 < right {
        mid := (left + right) / 2
        num := nums[mid]
    
        if num == target {
            return mid
        }

        if num <= nums[right] { // right is sorted
            if target <= nums[right] && target > num {
                left = mid
            }else {
                right = mid
            }
        }
        
        if num >= nums[left] {  // left is sorted
            if target >= nums[left] && target < num {
                right = mid
            }else {
                left = mid
            }
        }
    }

    return -1
}