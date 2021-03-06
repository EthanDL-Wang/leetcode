/***********************************************************************************************************
*执行用时：12 ms, 在所有 Go 提交中击败了52.73%的用户
*内存消耗：7.3 MB, 在所有 Go 提交中击败了5.095.09%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
************************************************************************************************************
*/

func firstUniqChar(s string) int {

    buff := make(map[rune]int, 0)
    var res int = -1

    for _, v := range s {
        buff[v]++
    }

    for i, v := range s {
        if buff[v] == 1 {
            res = i
            break
        }
    }

    return res
}