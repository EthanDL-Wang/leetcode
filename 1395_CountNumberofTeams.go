/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了55.07%的用户
*内存消耗：2.1 MB, 在所有 Go 提交中击败了25.00%的用户
*1.穷举
************************************************************************************************************
*/

/***********************************************************************************************************
n 名士兵站成一排。每个士兵都有一个 独一无二 的评分 rating 。

每 3 个士兵可以组成一个作战单位，分组规则如下：

从队伍中选出下标分别为 i、j、k 的 3 名士兵，他们的评分分别为 rating[i]、rating[j]、rating[k]
作战单位需满足： rating[i] < rating[j] < rating[k] 或者 rating[i] > rating[j] > rating[k] ，其中  0 <= i < j < k < n
请你返回按上述条件可以组建的作战单位数量。每个士兵都可以是多个作战单位的一部分。
************************************************************************************************************
*/
func numTeams(rating []int) int {

    lenRating := len(rating)
    if rating == nil || lenRating < 3 {
        return 0
    }

    var count int

    for i := 0; i < lenRating; i++ {
        for j := i+1; j < lenRating; j++ {
            if rating[i] < rating[j] {
                for k := j+1; k < lenRating; k++ {
                    if rating[j] < rating[k] {
                        count++
                    }
                }
            } else if rating[i] > rating[j] {
                for k := j+1; k < lenRating; k++ {
                    if rating[j] > rating[k] {
                        count++
                    }
                }
            }              
        }        
    }

    return count

}