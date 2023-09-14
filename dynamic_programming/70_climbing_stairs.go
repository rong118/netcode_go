func climbStairs(n int) int {
    if n == 1 { return 1}
    m := make([]int, n)
    m[0] = 1
    m[1] = 2
	
    for i := 2;i < n;i++ {
        m[i] = m[i - 1] + m[i -2]
    }

    return m[n - 1]
}