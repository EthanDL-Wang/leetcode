/***********************************************************************************************************
*执行用时：12 ms, 在所有 Go 提交中击败了43.11%的用户
*内存消耗：3.4 MB, 在所有 Go 提交中击败了84.56%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
************************************************************************************************************
*/

func findKthLargest(nums []int, k int) int {

    sort.Ints(nums)
    //fmt.Println(nums)
    return nums[len(nums)-k]



}