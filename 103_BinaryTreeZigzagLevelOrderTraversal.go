/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2.5 MB, 在所有 Go 提交中击败了65.47%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
************************************************************************************************************
*/



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {

    if root == nil {
        return nil
    }

    answer := make([][]int, 0)

    buff1 := make([]*TreeNode, 0)
    left := true
    buff1 = append(buff1, root)

    for len(buff1) > 0 {
        arrayTmp := make([]int, 0)
        if left {
            for i := 0; i < len(buff1); i++ {
                arrayTmp = append(arrayTmp, buff1[i].Val)
            }
            left = false
        } else {
            for i := len(buff1)-1; i >= 0; i-- {
                arrayTmp = append(arrayTmp, buff1[i].Val)
            }
            left = true
        }
        answer = append(answer, arrayTmp)
        


        buff2 := make([]*TreeNode, 0)
        for i := 0; i < len(buff1); i++ {
            if buff1[i].Left != nil {
                buff2 = append(buff2, buff1[i].Left)
            }
            
            if buff1[i].Right != nil {
                buff2 = append(buff2, buff1[i].Right)
            } 
            
        }

        buff1 = buff2


    }


    return answer

}