/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：3.9 MB, 在所有 Go 提交中击败了48.10%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。

如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。

请你返回排序后的数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


func getCount(num int) int {
    var count int 
    var tmp int = num

    for tmp > 0 {

        if tmp & 0x1 == 1 {
            count++
        }
        tmp >>= 1
    }

    return count
}

func sortByBits(arr []int) []int {

    arrLen := len(arr)

    buff := make([]int, arrLen)
    answer := make([]int, arrLen)

    for i, v := range arr {
        buff[i] = getCount(v)
    }

    var start int
    var count int 

    for i := 0; i < 16; i++ {

        count = 0
        for j := 0; j < arrLen; j++ {
            if buff[j] == i {
                answer[start+count] = arr[j]
                count++
            }
        }

        slice := answer[start:start+count]
        sort.Ints(slice)

        start += count
    }

    return answer

}