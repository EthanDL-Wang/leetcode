/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2.8 MB, 在所有 Go 提交中击败了100.00%的用户
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