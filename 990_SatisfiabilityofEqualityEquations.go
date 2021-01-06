/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2.8 MB, 在所有 Go 提交中击败了100.00%的用户
1.遍历equations，处理“==”的部分，buff为26个元素的数组，每个元素代表一个字母的权重
2.判断==两边的字母的权重，如果都==0，则计为index（初始为1），index++
3.如果一个字母>0,另一个字母==0，则代表等于0的那个字母第一次出现，权重与>0的字母相同，将两个字母权重改为相同
4.如果都>0， 就判断两个字母的权重是否相同，不同的时候，需要处理。遍历buff将所有改为相同
5.遍历遍历equations，处理“!=“的部分
6.如果！=两边的字母相同，直接返回false
7.如果！=两边的字母都>0，那权重相等即返回false

8.最终遍历结束返回true

作者：346998614
链接：https://leetcode-cn.com/problems/satisfiability-of-equality-equations/solution/shuang-bai-by-346998614-k6o5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。

只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/satisfiability-of-equality-equations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。                                                 
************************************************************************************************************
*/


func equationsPossible(equations []string) bool {

    var buff [26]int
    index := 1

    for _, v := range equations {

        if strings.Contains(v, "==") {
            if buff[v[0]-'a'] > 0 && buff[v[3]-'a'] > 0 {
                if buff[v[0]-'a'] != buff[v[3]-'a'] {
                    tmp := buff[v[0]-'a']
                    for i := 0; i < 26; i++ {
                        if buff[i] == tmp {
                            buff[i] = buff[v[3]-'a']
                        }
                    }
                }
            } else if buff[v[0]-'a'] == 0 && buff[v[3]-'a'] > 0 {
                buff[v[0]-'a'] = buff[v[3]-'a']
            } else if buff[v[0]-'a'] > 0 && buff[v[3]-'a'] == 0 {
                buff[v[3]-'a'] = buff[v[0]-'a']
            } else if buff[v[0]-'a'] == 0 && buff[v[3]-'a'] == 0 {
                buff[v[0]-'a'] = index
                buff[v[3]-'a'] = index
                index++
            }
        }
    }

    //fmt.Println(buff)

    for _, v := range equations {
        if strings.Contains(v, "!=") {
            if v[0] == v[3] {
                return false
            }

            if buff[v[0]-'a'] > 0 && buff[v[3]-'a'] > 0 {
                if buff[v[0]-'a'] == buff[v[3]-'a'] {
                    return false
                }
            }
        }
    }


    return true

}