func longestConsecutive(nums []int) int {
    m := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
        m[nums[i]] = true
    }

    var ans = 1;
    for _, num := range(nums) {
        if m[num - 1] {
            continue
        }

        l := 1;
        next := num + 1;

        for m[next] {
            l++;
            next++;
        }

        if l > ans {
            ans = l;
        }
    }

    return ans
}