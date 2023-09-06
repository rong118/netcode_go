func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)

    for _, v := range(nums) {
        m[v]++
    }

    // Create a slice to store key-value pairs
    var keyValuePairs []struct {
        Key   int
        Value int
    }

    // Populate the slice with the map data
    for key, value := range m {
        keyValuePairs = append(keyValuePairs, struct {
            Key   int
            Value int
        }{
            Key:   key,
            Value: value,
        })
    }

    // Sort slice based on struct value
    sort.SliceStable(keyValuePairs, func(i, j int) bool {
        return keyValuePairs[i].Value > keyValuePairs[j].Value
    })

    ans := make([]int, k)

    for i := 0; i < k; i++ {
        ans[i] = keyValuePairs[i].Key
    }

    return ans
}