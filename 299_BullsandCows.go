/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2.2 MB, 在所有 Go 提交中击败了88.10%的用户
*1.从后往前逐字匹配，查找bulls的个数
*2.将secret的元素，放入大小为10的数组，下标代表数字，元素大小代表数字的个数
*3.判断guess中的元素，有多少个secret中的元素相同，剪掉bulls的个数，即为cows的个数
************************************************************************************************************
*/

func getHint(secret string, guess string) string {

    var bullsCount, cowsCount int

    lenSecret := len(secret)
    lenGuess := len(guess)

    for i, j := lenSecret-1, lenGuess-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if secret[i] == guess[j] {
            bullsCount++
        }
    }

    slice := make([]int, 10)

    for _, v := range []byte(secret){
        slice[v-'0']++
    }

    for _, v := range []byte(guess){
        if slice[v-'0'] > 0 {
            cowsCount++
            slice[v-'0']--
        }
    }    

    cowsCount -= bullsCount

    var buff string

    buff = strconv.Itoa(bullsCount) + "A" + strconv.Itoa(cowsCount) + "B"

    return buff

}