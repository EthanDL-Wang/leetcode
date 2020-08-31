/***********************************************************************************************************
*执行用时：32 ms, 在所有 Go 提交中击败了82.52%的用户
*内存消耗：6.2  MB, 在所有 Go 提交中击败了95.56%的用户
*1.桶排序
************************************************************************************************************
*/


func findErrorNums(nums []int) []int {

    len := len(nums)

    buff := make([]int, len+1)
    answer := make([]int, 2)

    for _, v := range nums {
        buff[v]++
    }

    for i, v := range buff {

        if i == 0 {
            continue
        }

        if v == 2 {
            answer[0] = i
        }

        if v == 0 {
            answer[1] = i
        }
    }


    return answer
}