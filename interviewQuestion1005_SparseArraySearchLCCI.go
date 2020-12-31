/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了94.81%的用户
*内存消耗：5.7 MB, 在所有 Go 提交中击败了63.16%的用户
*1.words的元素作为key，下标作为value
*2.边放入map，边查找
************************************************************************************************************
*/

/***********************************************************************************************************
*稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
************************************************************************************************************
*/

func findString(words []string, s string) int {
    len := len(words)
    buffMap := make(map[string]int,len)

    var index int = -1

    for i, v := range words {
        if v == "" {
            continue
        } else {
            buffMap[v] = i
        }

        _, ok := buffMap[s] 
        if ok == true {
            index = buffMap[s]
            break
        }

    }


    return index
}