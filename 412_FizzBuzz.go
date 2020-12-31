/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：3.5 MB, 在所有 Go 提交中击败了73.42%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。
************************************************************************************************************
*/

func fizzBuzz(n int) []string {

    answer := make([]string, n)

    for i := 1; i <= n; i++ {
        if i % 5 == 0 && i % 3 == 0 {
            answer[i-1] = fmt.Sprintf("%s", "FizzBuzz")
        } else if i % 5 == 0 && i % 3 != 0 {
            answer[i-1] = fmt.Sprintf("%s", "Buzz")
        } else if i % 5 != 0 && i % 3 == 0 {
            answer[i-1] = fmt.Sprintf("%s", "Fizz")
        } else {
            answer[i-1] = fmt.Sprintf("%d", i)
        }
    }

    return answer

}