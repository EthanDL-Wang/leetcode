/***********************************************************************************************************
*解题思路：
*1.转圈跑，那么我们只关心开始和结束的点就好了
*2.从开始到结束，肯定是经过的最多的
*3.首先需要计算出从开始到结束的区域个数。
*4.然后依次将区域装进buff。
*5.最后将buff排序
*6.开始和结束相等时，为特殊情况。
************************************************************************************************************
*/

func mostVisited(n int, rounds []int) []int {

    len := len(rounds)
    start := rounds[0]
    end := rounds[len-1]

    var buffLen int

    if end == start {
        buffLen = 1
        buff := make([]int, 1)
        buff[0] = start
        return buff
    } else if end > start {
        buffLen = end - start +1
    } else if end < start {
        buffLen = n - start + end + 1
    }

    buff := make([]int, buffLen)

    for i := 0; i < buffLen; i++ {
        buff[i] = (start+i)%n
        if buff[i] == 0 {
            buff[i] = n
        }
    }

    sort.Ints(buff)

    return buff
}