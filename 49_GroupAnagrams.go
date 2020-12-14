/***********************************************************************************************************
*执行用时：16 ms, 在所有 Go 提交中击败了99.85%的用户
*内存消耗：7.1 MB, 在所有 Go 提交中击败了92.27%的用户
*1.将strs的元素按字符数组排序，作为map的key
*2.将key添加进map的次序作为value
*3.即value即为key第几次添加进map的
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
************************************************************************************************************
*/

func groupAnagrams(strs []string) [][]string {

    buff := make(map[string]int)
    index := 0

    res := make([][]string, 0)


    for _, v := range strs {

        tmpBuff := make([]int, 26)
        for _, v := range v {
            tmpBuff[v-'a']++
        }

        tmpKey := make([]byte, 0, len(v))
        for i := 0; i < 26; i++ {
            for j := 0; j < tmpBuff[i]; j++{
                tmpKey = append(tmpKey, byte(i)+'a')
            }
        }

        //fmt.Println(string(tmpKey))

        _, ok := buff[string(tmpKey)]
        if ok {
            res[buff[string(tmpKey)]] = append(res[buff[string(tmpKey)]], v)
        } else {
            buff[string(tmpKey)] = index
            index++
            tmpRes := make([]string, 0)
            tmpRes = append(tmpRes, v)
            res = append(res, tmpRes)
        }
    }

    return res

}