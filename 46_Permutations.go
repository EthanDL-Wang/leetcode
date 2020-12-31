/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2.7 MB, 在所有 Go 提交中击败了52.59%的用户
*1.解题思路：按区间开始大小排序后，再合并
*2.将区间开始端作为key，结束点作为value，放入map，目的是合并开始端一样的区间
*3.将map中的key放入数组left，排序
*4.按照map和left的数据，将区间结束点放入right。至此，left[0]到right[0]即为一个区间
*5.创建新的map，遍历left，遍历map，合并重合的区间
*6.新map的值即为答案
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
************************************************************************************************************
*/




var res [][]int

func dps(a []int, visit map[int]bool, nums []int){

    if len(nums) == len(a) {
        
        tmp := make([]int, len(nums))
        copy(tmp, a)
        res = append(res, tmp)
        return 
    }

    for _, v := range nums {
        if visit[v] == true {
            continue
        }

        a = append(a, v)
        visit[v] = true
        dps(a, visit, nums)
        a = a[:len(a)-1]
        visit[v] = false
    }



}



func permute(nums []int) [][]int {

    res = make([][]int, 0)
    visited := make(map[int]bool, len(nums))

    a := []int{}

    dps(a, visited, nums)

    return res

}