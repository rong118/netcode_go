func maxProfit(prices []int) int {
    min := math.MaxUint32
    ans := 0

    for _, price := range prices {
        if price < min {
            min = price
        }else{
            gain := price - min
            if gain > ans {
                ans = gain
            }
        }
    }

    return ans
}