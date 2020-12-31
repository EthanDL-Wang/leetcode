/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2 MB, 在所有 Go 提交中击败了59.39%的用户
*1.回溯
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
************************************************************************************************************
*/


var phoneNumber map[string]string = map[string]string{
    "2":"abc",
    "3":"def",
    "4":"ghi",
    "5":"jkl",
    "6":"mno",
    "7":"pqrs",
    "8":"tuv",
    "9":"wxyz",
}


var res []string

func backTracking(digits string, index int, tmp string) {
    if len(tmp) == len(digits) {
        res = append(res, string(tmp))
    } else {
        letters := phoneNumber[string(digits[index])]
        numberLen := len(letters)

        for i := 0; i < numberLen; i++ {
            backTracking(digits, index+1, tmp+string(letters[i]))
        }

    }
}


func letterCombinations(digits string) []string {
    res = make([]string, 0)

    if len(digits) <= 0 {
        return res
    }

    backTracking(digits, 0, "")

    return res 

}