/***********************************************************************************************************
*1.一个单词在其中一个句子中只出现一次，在另一个句子中却没有出现，那么这个单词就是不常见的。
*2.那么总共只出现一次的单词即为不常见的单词
*3.以单词作为map的key，出现的次数为value
*4.挑选只出现一次的词
************************************************************************************************************
*/

func uncommonFromSentences(A string, B string) []string {

    var count int
    var index int
    sliceA := strings.Split(A, " ")
    sliceB := strings.Split(B, " ")
    len := len(sliceA) + len(sliceB)
    buffMap := make(map[string]int, len)

    for _, v := range sliceA {
        buffMap[string(v)]++
    }

    for _, v := range sliceB {
        buffMap[string(v)]++
    }    
    
    for _, v := range buffMap {
        if v == 1 {
            count++
        }
    }

    buff := make([]string, count)

    for k, v := range buffMap {
        if v == 1 {
            buff[index] = k
            index++
        }
    }    

    return buff

}