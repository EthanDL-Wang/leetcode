/***********************************************************************************************************
*1.第一次map，将arr的元素作为key，个数作为value
*2.第二次map，将上次map的value作为key，个数作为value
*2.第二次map中的value如果都步大于1，则为true
*
*
************************************************************************************************************
*/


func uniqueOccurrences(arr []int) bool {

    len := len(arr)

    buffMap1 := make(map[int]int, len)
    buffMap2 := make(map[int]int, len)

    for i := 0; i < len; i++ {
        buffMap1[arr[i]]++
    }

    for _, v := range buffMap1 {        
        buffMap2[v]++

        if buffMap2[v] > 1 {
            return false
        }

    } 

    return true

}