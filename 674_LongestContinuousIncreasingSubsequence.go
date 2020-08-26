func getLength(nums []int, start int) int {
    len := len(nums)
    count := 1 

    tmp := nums[start]
    for i := start+1; i < len; i++ {
        if tmp < nums[i] {
            tmp = nums[i]
            count++
        } else {
            break
        }
    }

    return count
}

func findLengthOfLCIS(nums []int) int {

    if nums == nil || len(nums) <= 0 {
        return 0
    }

    len := len(nums)
    max := 1

    for i := 0; i < len; i++ {
        tmp := getLength(nums, i)
        if tmp > max {
            max = tmp
        }
        i += tmp-1
    }

    return max
}