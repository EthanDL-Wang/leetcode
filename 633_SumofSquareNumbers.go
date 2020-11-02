/***********************************************************************************************************
*执行用时：16 ms, 在所有 Go 提交中击败了17.81%的用户
*内存消耗：1.9 MB, 在所有 Go 提交中击败了84.72%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。
************************************************************************************************************
*/

func judgeSquareSum(c int) bool {

    var num int

    for num * num <= c {
        if int(math.Pow(math.Floor(math.Sqrt(float64(c-num*num))), 2)) == c - num*num {
            return true
        } 
        num++ 
    }

    return false

}