/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2.4 MB, 在所有 Go 提交中击败了5.56%的用户
*1.找字串，然后作为key，放入map，并记录个数
*2.将map的key，放入buff，即为答案
************************************************************************************************************
*/

/***********************************************************************************************************
给你一个字符串数组 words ，数组中的每个字符串都可以看作是一个单词。请你按 任意 顺序返回 words 中是其他单词的子字符串的所有单词。

如果你可以删除 words[j] 最左侧和/或最右侧的若干字符得到 word[i] ，那么字符串 words[i] 就是 words[j] 的一个子字符串。

************************************************************************************************************
*/

func stringMatching(words []string) []string {

    lenWords := len(words)
    buffMap := make(map[string]bool, lenWords)
    var count int 

    for i := 0; i < lenWords; i++ {
        for j := 0; j < lenWords; j++ {
            if i == j {
                continue
            }

            if len(words[i]) > len(words[j]) {
                if strings.Index(words[i], words[j]) != -1 {
                    _, ok := buffMap[words[j]]
                    if ok == false {
                        buffMap[words[j]] = true
                        count++
                    }
                }
            }

        }
    }

    buff := make([]string, count)
    count = 0

    for k ,_ := range buffMap {
        buff[count] = k
        count++
    }    

    return buff

}