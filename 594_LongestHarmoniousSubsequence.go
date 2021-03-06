/***********************************************************************************************************
*执行用时：72 ms, 在所有 Go 提交中击败了47.95%的用户
*内存消耗：6.6 MB, 在所有 Go 提交中击败了100.00%的用户
*1.和谐数组是指一个数组里元素的最大值和最小值之间的差别正好是1, 并不是需要连续。
*2.将元素作为key，出现次数作为value
*3.遍历map，计算key和key+1的value的总和，记录最大值
************************************************************************************************************
*/

func findLHS(nums []int) int {

    len := len(nums)

    buff := make(map[int]int,len)
    for i := 0; i < len; i++ {
        buff[nums[i]]++
    }

    var max int 

    for k,v := range buff {
        _, ok := buff[k+1]
        if ok == true {
            if max < v + buff[k+1] {
                max = v + buff[k+1]
            }
        }
    }

    return max
}
