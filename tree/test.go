package main

import "fmt"

func main() {
	var ans [][]int
	fmt.Println("Original ans:", ans)
    // Create a slice
    originalSlice := []int{1, 2, 3, 4, 5}
	ans = append(ans, originalSlice)
	fmt.Println("Original ans:", ans)

    // Clean the slice by truncating it
    originalSlice = originalSlice[:0]

    fmt.Println("Original ans:", ans)
	originalSlice = []int{9,8,7}
	fmt.Println("Original ans:", ans)
}
