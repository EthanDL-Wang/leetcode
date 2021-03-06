/***********************************************************************************************************
*执行用时：744 ms, 在所有 Go 提交中击败了26.21%的用户
*内存消耗：2.2 MB, 在所有 Go 提交中击败了93.29%的用户
*1.递归
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


var count int

func depth(nums []int, S int, index int, sum int) {
    if index == len(nums) {
        if sum == S {
            count++
        }
        return 
    } else {
        depth(nums, S, index+1, sum+nums[index])
        depth(nums, S, index+1, sum-nums[index])
    }
}

func findTargetSumWays(nums []int, S int) int {

    count = 0
    depth(nums, S, 0, 0)

    return count

}