/***********************************************************************************************************
*执行用时：48 ms, 在所有 Go 提交中击败了7.09%的用户
*内存消耗：7.5 MB, 在所有 Go 提交中击败了26.80%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。

请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/

func reconstructQueue(people [][]int) [][]int {

    buffMap := make(map[int]int, len(people))
    answer := make([][]int, 0, len(people))

    for _, v := range people {
        buffMap[v[0]]++
    }

    height := make([]int, 0)

    for k, v := range buffMap {
        if v > 0 {
            height = append(height, k)
        }
    }

    sort.Ints(height)
    for i := len(height)-1; i >= 0; i-- {
        heightCount := make([]int, 0)
        for j := 0; j < len(people); j++ {
            if people[j][0] == height[i] {    
                heightCount = append(heightCount, people[j][1])              
            }
        }

        sort.Ints(heightCount)
        for k := 0; k < len(heightCount); k++ {
            answer = append(answer[:heightCount[k]], append([][]int{[]int{height[i], heightCount[k]}}, answer[heightCount[k]:]...)...)
        }
                


    }

    return answer

}