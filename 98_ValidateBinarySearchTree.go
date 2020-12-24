/***********************************************************************************************************
*执行用时：8 ms, 在所有 Go 提交中击败了71.98%的用户
*内存消耗：5.4 MB, 在所有 Go 提交中击败了92.40%的用户
*1.回溯
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

 const (
    MAX_INT64 = 0xffffffffffffffff>>1
    MIN_INT64 = ^MAX_INT64
)


func reverse(node *TreeNode, left int, right int) bool {

    if node == nil {
        return true
    }

    if node.Val >= right || node.Val <= left {
        return false
    }

    return reverse(node.Left, left, node.Val) && reverse(node.Right, node.Val, right)



}


func isValidBST(root *TreeNode) bool {


    return reverse(root, MIN_INT64, MAX_INT64)

}