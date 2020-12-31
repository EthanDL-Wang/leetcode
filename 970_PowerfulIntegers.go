/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：3.1 MB, 在所有 Go 提交中击败了20.00%的用户
*1.创建数组，存储x,y的次方值
*2.计算x，y次方值的和，符合要求放入map
************************************************************************************************************
*/

/***********************************************************************************************************
*给定两个正整数 x 和 y，如果某一整数等于 x^i + y^j，其中整数 i >= 0 且 j >= 0，那么我们认为该整数是一个强整数。   
*返回值小于或等于 bound 的所有强整数组成的列表                                                                
*你可以按任何顺序返回答案。在你的回答中，每个值最多出现一次。                                                 
************************************************************************************************************
*/


func powerfulIntegers(x int, y int, bound int) []int {
    const MAX_LEN int = 21
    powX := make([]int, MAX_LEN)
    powY := make([]int, MAX_LEN)

    for i := 0; i < MAX_LEN; i++ {
        powX[i] = int(math.Pow(float64(x), float64(i)))
        powY[i] = int(math.Pow(float64(y), float64(i)))
    }

    //fmt.Println(powX,powY)

    buffMap := make(map[int]bool, MAX_LEN*MAX_LEN)
    var count int 
    for i := 0; i < MAX_LEN; i++ {

        if powX[i] >= bound {
            break
        }

        for j := 0; j < MAX_LEN; j++ {
            if powX[i]+powY[j] > bound {
                break
            }

            _, ok := buffMap[powX[i]+powY[j]]
            if ok == false {
                buffMap[powX[i]+powY[j]] = true
                count++
            }

        }
        //fmt.Println(buffMap)
    }

    answer := make([]int, count)
    count=0
    for k, _ := range buffMap {
        answer[count] = k
        count++
    }

    return answer

}