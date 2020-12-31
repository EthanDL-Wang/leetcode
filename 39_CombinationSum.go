/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了78.03%的用户
*内存消耗：3.2 MB, 在所有 Go 提交中击败了40.27%的用户
*func dfs(target int, candidates []int, combine []int, index int)
*index:去重
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。                                              
************************************************************************************************************
*/

var res [][]int

func dfs(target int, candidates []int, combine []int, index int) {


    if target == 0 {
        tmp := make([]int, len(combine))
        copy(tmp, combine)
        res = append(res, tmp)
        return
    } else if target < 0 {
        return
    }

    for i := index; i < len(candidates); i++ {
        combine = append(combine, candidates[i])
        dfs(target-candidates[i], candidates, combine, i)
        combine = combine[:len(combine)-1]
    }

}


func combinationSum(candidates []int, target int) [][]int {

    res = make([][]int, 0)

    combine := make([]int, 0)

    dfs(target, candidates, combine, 0)

    return res

}