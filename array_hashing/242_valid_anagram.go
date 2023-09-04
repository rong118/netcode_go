func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := make(map[byte]int)
    for idx := 0; idx < len(s); idx++ {
        m[s[idx] - 'a']++
        m[t[idx] - 'a']--
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}