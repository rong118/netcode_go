func lengthOfLongestSubstring(s string) int {
    // Convert the string to a slice of runes
    runes := []rune(s)

    l := 0
    r := 0
    m := make(map[rune]int) //rune, idx
    ans := 1

    for r < len(runes) {
        _, found := m[runes[r]]
        if found {
            l = m[runes[r]] + 1
        }
        
        m[runes[r]] = r
        str_len := r - l + 1
        if str_len > ans {
            ans = str_len
        }
        r++;
    }

    return ans
}