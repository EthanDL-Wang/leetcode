/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：2.1 MB, 在所有 Go 提交中击败了85.71%的用户
*1.回溯
************************************************************************************************************
*/

/***********************************************************************************************************
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
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

 func reverse(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if (p != nil && q == nil) || (p == nil && q != nil) {
        return false
    } else if p != nil && q != nil {
        if p.Val != q.Val {
            return false
        }
    }

    return reverse(p.Left, q.Left) && reverse(p.Right, q.Right)


}


func isSameTree(p *TreeNode, q *TreeNode) bool {

    return reverse(p, q)

}