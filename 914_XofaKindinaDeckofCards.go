
/***********************************************************************************************************
*执行用时：16 ms, 在所有 Go 提交中击败了90.00%的用户
*内存消耗：6.1 MB, 在所有 Go 提交中击败了10.00%的用户
*1.
************************************************************************************************************
*/

/***********************************************************************************************************
给定一副牌，每张牌上都写着一个整数。

此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：

每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


func hasGroupsSizeX(deck []int) bool {

    buff := make(map[int]int, len(deck))

    for _, v := range deck {
        buff[v] += 1
    }

    countBuff := make([]int, 0)

    for _, v := range buff {
        if v == 1 {
            return false
        }
        countBuff = append(countBuff, v)
    }

    sort.Ints(countBuff)

    for i := 2; i < countBuff[0]; i++ {
        j := 0
        for j = 0; j < len(countBuff); j++ {
            if countBuff[j]%i != 0 {
                break
            }
        }
        if j == len(countBuff) {
            return true
        }

    }     


    for i := 0; i < len(countBuff); i++ {
        j := 0
        for j = 0; j < len(countBuff); j++ {
            if countBuff[j]%countBuff[i] != 0 {
                break
            }
        }
        if j == len(countBuff) {
            return true
        }

    } 

    return false

}