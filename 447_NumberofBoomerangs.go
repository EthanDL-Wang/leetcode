/***********************************************************************************************************
*执行用时：248 ms, 在所有 Go 提交中击败了14.85%的用户
*内存消耗：6.7 MB, 在所有 Go 提交中击败了56.25%的用户
*1.选定points中的一个点，计算其余点到这一点的距离，将距离作为key存储
*2.遍历整个map，value>=2时，及可组成回旋镖，个数为value*(value-1)
*3.将第一点和第二点的操作，遍历整个points
*4.相加得到返回值
************************************************************************************************************
*/


func numberOfBoomerangs(points [][]int) int {

    len := len(points)
    if points == nil || len < 3 {
        return 0
    }

    var count int 

    for i := 0; i < len; i++ {
        buffMap := make(map[int]int,len)
        for j := 0; j < len; j++ {
            key := int(math.Pow(math.Abs(float64(points[i][0]-points[j][0])), 2) + math.Pow(math.Abs(float64(points[i][1]-points[j][1])), 2))
            buffMap[key]++
        }

        for _, v := range buffMap {
            if v >= 2 {
                count += v*(v-1)
            }
        }

    }

    return count

}