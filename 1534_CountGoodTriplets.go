func countGoodTriplets(arr []int, a int, b int, c int) int {

    len := len(arr)
    var count int

    for i := 0; i < len - 2; i++ {
        for j := i+1; j < len - 1; j++ {

            if math.Abs(float64(arr[i] - arr[j])) > float64(a) {
                continue
            } 

            for k := j+1; k < len; k++ {
                if math.Abs(float64(arr[j] - arr[k])) > float64(b) {
                    continue
                } else {
                    if math.Abs(float64(arr[i] - arr[k])) <= float64(c) {
                        count++
                    }
                }
            }
        }
    }

    return count

}