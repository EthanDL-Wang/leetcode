/***********************************************************************************************************
*1.将text放入map
*2.计算balloon字符的个数，将l,o减半
*3.求最小值
************************************************************************************************************
*/

func maxNumberOfBalloons(text string) int {

    const LEN int = 5
    buff := []byte(text)
    buffMap := make(map[byte]int, LEN)



    var count int = 10000
    balloon := [...]byte{'b','a','l','o','n'}

    for _, v := range balloon {
        buffMap[v] = 0
    }

    for _, v := range buff {
        _, ok := buffMap[v] 
        if ok == true {
            buffMap[v]++
        }
    }

    buffMap['l'] /= 2
    buffMap['o'] /= 2

    for _, v := range buffMap {
        if v < count {
            count = v
        }
    }

    return count

}