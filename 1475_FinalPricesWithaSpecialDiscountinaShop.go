func finalPrices(prices []int) []int {

    len := len(prices)

    for i := 0; i < len; i++ {
        for j := i+1; j < len; j++ {
            if prices[i] >= prices[j] {
                prices[i] -= prices[j]
                break 
            }
        }
    }

    return prices
}