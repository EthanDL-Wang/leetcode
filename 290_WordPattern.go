/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2 MB, 在所有 Go 提交中击败了44.90%的用户
*1.将pattern的元素作为key，str中的单词作为value，放入map
*2.只是需要双向验证，key存在时，就要判断对应的value与当前str中的单词是否相同
*3.key不存在时，就需要判断当前str的单词是否已经在value里面
************************************************************************************************************
*/

func wordPattern(pattern string, str string) bool {
    
    lenPattern := len(pattern) 
    sliceStr := strings.Split(str," ")
    lenSliceStr := len(sliceStr)

    if lenPattern != lenSliceStr {
        return false
    }

    buffMap := make(map[byte]string, lenPattern)

    for i, v := range []byte(pattern) {
        _, ok := buffMap[v]
        if ok == true {
            if strings.Compare(sliceStr[i], buffMap[v]) != 0 {
                return false
            } 
        } else {

            for _, value := range buffMap {
                if strings.Compare(value, sliceStr[i]) == 0{
                    return false
                }
            }

            buffMap[v] = sliceStr[i]
        }
    }

    return true

}