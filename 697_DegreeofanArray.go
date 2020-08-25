/**********************************************************************************************************
*1.将nums的元素作为key，数量作为value，放入buffMap的map里面
*2.找到nums最大的度
*3.记录最大的度对应的nums的值
*4.计算nums值中，最小的度
***********************************************************************************************************
*/



func findShortestSubArray(nums []int) int {

    var maxDegree int
    var count int
    var index int
    var start, end int 
    var value int

    len := len(nums)
    buffMap := make(map[int]int, len)

    for i := 0; i < len; i++ {
        buffMap[nums[i]]++
    }

    value = buffMap[nums[0]]
    for _, v := range buffMap {
        if v > value {
            value = v
        }
    }

    for _, v := range buffMap {
        if v == value {
            count++
        }
    }

    
    buffKey := make([]int, count)
    for k, v := range buffMap {
        if v == value {
            buffKey[index] = k
            index++
        }
    }


    maxDegree = len-1
    for i := 0; i < count; i++ {
        start = 0
        end = len-1

        for j := 0; j < len; j++ {
            if nums[j] == buffKey[i] {
                start = j
                break
            }
        }    

        for j := len-1; j >= 0; j-- {
            if nums[j] == buffKey[i] {
                end = j
                break
            }
        }

        if maxDegree > end - start {
           maxDegree = end - start 
        }

    }

    return maxDegree+1

}