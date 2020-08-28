/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了92.22%的用户
*内存消耗：3.5 MB, 在所有 Go 提交中击败了43.48%的用户
*1.定义两个变量，updown，leftright
*2.U时updown++，D时--
*3.R时leftright++，L时--
*4.最后判断updown，leftright都为0时，返回true，否则为false
************************************************************************************************************
*/

func judgeCircle(moves string) bool {

    var upDown int
    var leftRight int

    for _, v := range moves {
        if v == 'U' {
            upDown++
        } else if v == 'D' {
            upDown--
        } else if v == 'R' {
            leftRight++
        } else if v == 'L' {
            leftRight--
        }
    }

    if leftRight == 0 && upDown == 0 {
        return true 
    }

    return false
}