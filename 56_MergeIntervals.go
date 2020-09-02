/***********************************************************************************************************
*执行用时：12 ms, 在所有 Go 提交中击败了83.70%的用户
*内存消耗：5.1 MB, 在所有 Go 提交中击败了13.63%的用户
*1.解题思路：按区间开始大小排序后，再合并
*2.将区间开始端作为key，结束点作为value，放入map，目的是合并开始端一样的区间
*3.将map中的key放入数组left，排序
*4.按照map和left的数据，将区间结束点放入right。至此，left[0]到right[0]即为一个区间
*5.创建新的map，遍历left，遍历map，合并重合的区间
*6.新map的值即为答案
************************************************************************************************************
*/

/***********************************************************************************************************
给出一个区间的集合，请合并所有重叠的区间。
************************************************************************************************************
*/

func merge(intervals [][]int) [][]int {

    lenIntervals := len(intervals)
    buffMap := make(map[int]int, lenIntervals)
    var buffMapCount int

    for _, v := range intervals {
        _, ok := buffMap[v[0]]
        if ok == true {
            if buffMap[v[0]] < v[1] {
                buffMap[v[0]] = v[1]
            }
        } else {
            buffMap[v[0]] = v[1]
            buffMapCount++
        }   
    }     
    
    buffLeft := make([]int, buffMapCount)
    var index int 
    for k, _ := range buffMap {
        buffLeft[index] = k
        index++
    }

    sort.Ints(buffLeft)
    buffRight := make([]int, buffMapCount)

    index = 0
    for i := 0; i < buffMapCount; i++ {
        buffRight[index] = buffMap[buffLeft[i]]
        index++
    }

    newBuffMap := make(map[int]int, buffMapCount)

    flag := 0
    for i := 0; i < buffMapCount; i++ {
        flag = 0
        for k, v := range newBuffMap {           
            if buffLeft[i] <= v {
                if buffRight[i] > v {
                    newBuffMap[k] = buffRight[i]                   
                }
                flag = 1
                break
            }      
        }
        if flag == 0 {
            newBuffMap[buffLeft[i]] = buffRight[i]
        }
        
    }


    index = 0
    for k, _ := range newBuffMap {
        _, ok := newBuffMap[k]
        if ok == true {
            index++
        }
        
    }

    //fmt.Printf("index=%d\n",index)

    answer := make([][]int, index)
    for i := 0; i < index; i++ {
        answer[i] = make([]int,2)
    }

    index = 0
    for k, v := range newBuffMap {
        answer[index][0] = k
        answer[index][1] = v
        index++
    }    

    
    return answer



}