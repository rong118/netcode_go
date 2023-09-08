func isValid(s string) bool {
    var stk []rune
    // Convert the string to a slice of runes
    runes := []rune(s)

    for _, v := range(runes) {
        if v == '(' || v == '{' || v == '[' {
            stk = append(stk, v)
            continue
        }

        n := len(stk)
        if n == 0 { return false }
        if non_match(stk[n - 1], v) {
            return false
        }
        stk = stk[:n - 1]
    }

    return len(stk) == 0
}

func non_match(a rune, b rune) bool{
    if a == '(' && b == ')' {
        return false
    }else if a == '{' && b == '}' {
        return false
    }else if a == '[' && b == ']' {
        return false
    }

    return true
}