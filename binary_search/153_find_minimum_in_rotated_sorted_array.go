func findMin(nums []int) int {
    left, right := 0, len(nums) - 1
    
    for left + 1 < right {
        mid := (left + right) / 2

        // Left sorted portion
        if nums[left] <= nums[mid]{
           if nums[left] <= nums[right] {
               right = mid
           }else {
               left = mid
           }
        // Right sorted portion
        } else {
            right = mid
        }
    }
    if nums[left] < nums[right]{
        return nums[left]
    }else {
        return nums[right]
    }
}