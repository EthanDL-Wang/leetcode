/***********************************************************************************************************
*执行用时：8 ms, 在所有 Go 提交中击败了81.25%的用户
*内存消耗：5 MB, 在所有 Go 提交中击败了98.75%的用户
*1.桶排序，比较
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：

字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


func findAnagrams(s string, p string) []int {

    if len(s) <= 0 || len(s) < len(p) {
        return nil
    }


    res := make([]int, 0)

    var buffP [26]int
    var buffS [26]int

    for _, v := range p {
        buffP[v-'a']++
    }

    for i := 0; i < len(p); i++ {
        buffS[s[i]-'a']++
    }
    //fmt.Println(buffP)
    if buffP == buffS {
        
        //fmt.Println(buffS)
        res = append(res, 0)
    }

    for i := 1; i < len(s)-len(p)+1; i++ {


        buffS[s[i-1] - 'a']--
        buffS[s[i+len(p)-1] - 'a']++
        //fmt.Println(buffS)
        if buffP == buffS {
            res = append(res, i)
        }

    }


    return res
}