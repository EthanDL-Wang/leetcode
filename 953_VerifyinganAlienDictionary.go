/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
*内存消耗：2.6 MB, 在所有 Go 提交中击败了73.33%的用户
*1.将order放入map，元素作为key，下标作为value
*2.words中相邻的两个元素一次比较
************************************************************************************************************
*/

func isAlienSorted(words []string, order string) bool {

    const LETTER_LEN int = 26 

    buffLetter := make(map[byte]int, LETTER_LEN)

    for i := 0; i < LETTER_LEN; i++ {
        buffLetter[order[i]] = i
    }

    lenWords := len(words)
    //fmt.Printf("%T len=%d\n",words[0], len(words[0])) 
    
    len1 := len(words[0])
    for i := 0; i < lenWords-1; i++ {
        len2 := len(words[i+1])
        var index int
        for index = 0; index < len1 && index < len2; index++ {           
            if buffLetter[words[i][index]] != buffLetter[words[i+1][index]] {
                break
            }
        }
        //fmt.Println(index)

        if index < len1 && index < len2 {
            if buffLetter[words[i][index]] > buffLetter[words[i+1][index]] {
                return false
            }     
        } 
        
        if index == len1 || index == len2 {
            if buffLetter[words[i][index-1]] == buffLetter[words[i+1][index-1]]  && len1 > len2 {
                return false
            }
        } 

        len1 = len2
    }
    
    return true
}