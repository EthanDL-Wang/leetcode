/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：1.9 MB, 在所有 Go 提交中击败了21.61%的用户
*1.桶排序，比较
************************************************************************************************************
*/

/***********************************************************************************************************
小扣在秋日市集发现了一款速算机器人。店家对机器人说出两个数字（记作 x 和 y），请小扣说出计算指令：

"A" 运算：使 x = 2 * x + y；
"B" 运算：使 y = 2 * y + x。
在本次游戏中，店家说出的数字为 x = 1 和 y = 0，小扣说出的计算指令记作仅由大写字母 A、B 组成的字符串 s，字符串中字符的顺序表示计算顺序，请返回最终 x 与 y 的和为多少。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/nGK0Fy
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/

func calculate(s string) int {

    x := 1
    y := 0

    for _, v := range s {
        if v == 'A' {
            x = x * 2 + y
        } else {
            y = y * 2 + x
        }
    }

    return x + y

}