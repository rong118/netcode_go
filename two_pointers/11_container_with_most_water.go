func maxArea(height []int) int {
    n := len(height)
    l := 0
    r := n - 1
    var ans = 0
    for l < r {
        var h = 0
        if height[l] > height[r]{
            h = height[r]
        }else{
            h = height[l]
        }
        var area = h * (r - l)
        
        if area > ans {
            ans = area
        }

        if height[l] > height[r] {
            r--
        }else {
            l++
        }
    }

    return ans
}