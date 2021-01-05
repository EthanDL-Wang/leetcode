/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2 MB, 在所有 Go 提交中击败了83.96%的用户
*解题思路：使用栈
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


func decodeString(s string) string {

    stack := make([]byte, 0)
    stackIndex := -1

    res := make([]byte, 0)

    
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9'{
            tmpIndex := 0
            for tmpIndex = i; tmpIndex < len(s); tmpIndex++ {
                //fmt.Printf("%s %s\n",string(stack), string(res))
                if s[tmpIndex] == ']' {
                    leftIndex := strings.LastIndex((string)(stack), "[")
                    subStr := stack[leftIndex+1:stackIndex+1]
                    numIndex := 0
                    for numIndex = leftIndex-1; numIndex >= 0; i-- {
                        if stack[numIndex] >= '0' && stack[numIndex] <= '9' {
                            numIndex--
                        } else {
                            break
                        }
                    }

                    //fmt.Printf("numIndex=%d, subStr=[%s]\n", numIndex, subStr)

                    stackIndex = numIndex
                    stack = stack[:stackIndex+1]

                    if numIndex == -1 {
                        subNum, _ := strconv.Atoi(string(stack[:leftIndex]))
                        tmpRepeatStr := strings.Repeat(string(subStr), subNum)
                        res = append(res, []byte(tmpRepeatStr)...)
                        stack = stack[:0]
                        stackIndex = -1
                        break

                    } else {
                        subNum, _ := strconv.Atoi(string(stack[numIndex+1:leftIndex]))
                        tmpRepeatStr := strings.Repeat(string(subStr), subNum)
                        stack = append(stack, tmpRepeatStr...)
                        stackIndex += len(tmpRepeatStr)
                    }



                } else {
                    stack = append(stack, s[tmpIndex])
                    stackIndex++
                }
            }
            i = tmpIndex
        } else {
            res = append(res, s[i])
        }
        

        
    }

    return string(res)


}