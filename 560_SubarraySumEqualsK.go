/***********************************************************************************************************
*执行用时：356 ms, 在所有 Go 提交中击败了29.76%的用户
*内存消耗：6.1 MB, 在所有 Go 提交中击败了56.25%的用户
*1.穷举
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
************************************************************************************************************
*/







func subarraySum(nums []int, k int) int {

    lenNums := len(nums)
    var count int 

    for i := 0; i < lenNums; i++ {
        tmpSum := 0
        for j := i; j < lenNums; j++ {
            tmpSum += nums[j]
            if k == tmpSum {
                count++
            }
        }         
    } 

    return count

}