
/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了75.11%的用户
*内存消耗：2.9 MB, 在所有 Go 提交中击败了87.11%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个二叉树，检查它是否是镜像对称的。
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

 func dfs(left *TreeNode, right *TreeNode) bool {

    if left == nil && right == nil {
        return true
    } else if left != nil && right == nil {
        return false 
    } else if left == nil && right != nil {
        return false 
    }

    return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)

}


func isSymmetric(root *TreeNode) bool {
    
    if root == nil {
        return true
    }


    return dfs(root.Left, root.Right)
}