
/***********************************************************************************************************
*执行用时：8 ms, 在所有 Go 提交中击败了41.67%的用户
*内存消耗：3 MB, 在所有 Go 提交中击败了94.62%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。

如果数组元素个数小于 2，则返回 0。
************************************************************************************************************
*/


func maximumGap(nums []int) int {

    if nums == nil || len(nums) < 1 {
        return 0
    }

    sort.Ints(nums)
    result := 0



    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i-1] > result {
            result = nums[i] - nums[i-1]
        }
    }

    return result

}

