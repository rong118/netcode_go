func uniquePaths(m int, n int) int {
    // Create a 2D slice
    // This initializes a slice of slices with nil values.
    // You'll need to initialize each individual slice in a loop or using other methods.
    dp := make([][]int, m)

    // Initialize each row slice
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for x := 0; x < m; x++ { dp[x][0] = 1 }
    for y := 0; y < n; y++ { dp[0][y] = 1 }

    for x := 1; x < m; x++ {
        for y := 1; y < n; y++ {
            dp[x][y] = dp[x - 1][y] + dp[x][y - 1]
        }
    }

    return dp[m - 1][n - 1]
}