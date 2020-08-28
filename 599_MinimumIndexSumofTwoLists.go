/***********************************************************************************************************
*执行用时：32 ms, 在所有 Go 提交中击败了85.83%的用户
*内存消耗：6.5 MB, 在所有 Go 提交中击败了91.67%的用户
*1.将list1放入map，元素作为key，下标作为value
*2.设定最小索引为len1+len2
*3.将list2中的元素去map中匹配，匹配到以后，计算最小索引，与当前最小索引比较
*4.记录最小索引，并记录个数
*5.返回最小索引中所有值
************************************************************************************************************
*/

func findRestaurant(list1 []string, list2 []string) []string {

    len1 := len(list1)
    len2 := len(list2)
    buffMap := make(map[string]int, len1)

    var count int
    var minIdex int = len1+len2

    for i, v := range list1 {
        buffMap[v] = i
    }

    for i, v := range list2 {
        _, ok := buffMap[v]
        if ok == true {
            if minIdex > i + buffMap[v] {
                minIdex = i + buffMap[v]
                count = 1
            } else if minIdex == i + buffMap[v] {
                count++
            }
            
        }
    }    
    //fmt.Println(count, minIdex)


    buff := make([]string, count)
    count = 0

    for i, v := range list2 {
        _, ok := buffMap[v]
        if ok == true {
            if minIdex == i + buffMap[v] {
                buff[count] = v
                count++
            } 
            
        }
    }    


    return buff
}