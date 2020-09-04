/***********************************************************************************************************
*执行用时：144 ms, 在所有 Go 提交中击败了19.23%的用户
*内存消耗：14.4 MB, 在所有 Go 提交中击败了38.46%的用户
*1.把nums中的元素作为key，个数作为value
*2.将key放入数组，计算答案的个数。
*3.申请answer作为返回值。
************************************************************************************************************
*/

/***********************************************************************************************************
设计一个算法，找出数组中两数之和为指定值的所有整数对。一个数只能属于一个数对。
************************************************************************************************************
*/

func pairSums(nums []int, target int) [][]int {

    lenNums := len(nums)
    buffMap := make(map[int]int, lenNums)
    var count int


    for i := 0; i < lenNums; i++ {
        buffMap[nums[i]]++
    }

    keyCount := 0
    for k, _ := range buffMap {
        _,ok := buffMap[k]
        if ok == true {
            keyCount++
        }
    }

    buffKey := make([]int, keyCount)
    index := 0
    for k, _ := range buffMap {
        buffKey[index] = k
        index++
    }

    //fmt.Println(buffMap)
    //fmt.Println(buffKey)

    for i := 0; i < keyCount; i++ {
        _, ok := buffMap[target-buffKey[i]]
        if ok == true {
            if target == buffKey[i]*2  {
                count += buffMap[buffKey[i]]/2
            } else {
                if buffMap[buffKey[i]] > buffMap[target-buffKey[i]] {
                    count += buffMap[target-buffKey[i]]
                } else {
                    count += buffMap[buffKey[i]]
                }

            }
            delete(buffMap, buffKey[i])
        }
        //fmt.Println(count)
    }   

    //fmt.Println(count)

    answer := make([][]int, count)
    for i := 0; i < count; i++ {
        answer[i] = make([]int,2)
    }

    
    newBuffMap := make(map[int]int, lenNums)
    for i := 0; i < lenNums; i++ {
        newBuffMap[nums[i]]++
    }

    count = 0
    for i := 0; i < keyCount; i++ {
        _, ok := newBuffMap[target - buffKey[i]]
        if ok == true {
            if target == buffKey[i]*2  {
                j := count
                count += newBuffMap[buffKey[i]]/2
                for ; j < count; j++ {
                    answer[j][0] = buffKey[i]
                    answer[j][1] = buffKey[i]
                }
            } else {
                j := count
                if newBuffMap[buffKey[i]] > newBuffMap[target-buffKey[i]] {
                    count += newBuffMap[target-buffKey[i]]
                } else {
                    count += newBuffMap[buffKey[i]]
                }
                //fmt.Println(j, count)
                for ; j < count; j++ {
                    answer[j][0] = buffKey[i]
                    answer[j][1] = target - buffKey[i]
                }

            }
            delete(newBuffMap, buffKey[i])
        }
    }  

    return answer

}