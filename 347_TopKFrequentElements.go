/***********************************************************************************************************
*执行用时：24 ms, 在所有 Go 提交中击败了12.46%的用户
*内存消耗：6.9 MB, 在所有 Go 提交中击败了5.02%的用户
*1.首先k大于两倍的lenarr时，即返回arr中最大值即可
*2.否则，将挑战者放进一个队列。与winner进行比较
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
************************************************************************************************************
*/





func topKFrequent(nums []int, k int) []int {

    lenNums := len(nums)

    buffMap := make(map[int]int, lenNums)

    for i := 0; i < lenNums; i++ {
        buffMap[nums[i]]++
    }
    //fmt.Println(buffMap)
    buffCount := make([]int, lenNums)

    index := 0
    for _, v := range buffMap {
        buffCount[index] = v
        index++
    }

    sort.Ints(buffCount)
    //fmt.Println(buffCount)
    answer := make([]int, k)

    index = 0
    for i := lenNums-1; i >= lenNums - k ; i-- {
        for k, v := range buffMap {
            if v == buffCount[i] {
                answer[index] = k
                index++
                delete(buffMap, k)
                break
            }
            
        }
    }

    return answer

}