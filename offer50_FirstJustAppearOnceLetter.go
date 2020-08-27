/***********************************************************************************************************
*1.创建结构体letterInfo，记录每一个字母出现的次数以及最小下标
*2.创建letterInfo的数组，遍历s
*3.遍历letterInfo的数组，找出只出现一次的字母的最小下标
************************************************************************************************************
*/



func firstUniqChar(s string) byte {

    if s == "" || len(s) <= 0 {
        return ' '
    }

    type letterInfo struct {
        minIndex int
        count int
    }

    len := len(s)
    const LEN int = 26
    var min int = len
    buff := make([]letterInfo, LEN)

    
    //fmt.Println(len)


    for i := 0; i < LEN; i++ {
        buff[i].minIndex = len
    }

    for i, v := range []byte(s) {
        if buff[v - 'a'].minIndex > i {
            buff[v - 'a'].minIndex = i
        }

        buff[v - 'a'].count++
    }

    for i := 0; i < LEN; i++ {
        if buff[i].count == 1 {
            if buff[i].minIndex < min {
                min = buff[i].minIndex
            }            
        }
    }

    if min == len {
        return ' '
    }

    return byte(s[min])







}