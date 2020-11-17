/***********************************************************************************************************
*执行用时：24 ms, 在所有 Go 提交中击败了51.90%的用户
*内存消耗：6.8 MB, 在所有 Go 提交中击败了23.80%的用户
*func dfs(target int, candidates []int, combine []int, index int)
*index:去重
************************************************************************************************************
*/

/***********************************************************************************************************
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-binary-trees
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
 func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

    if t1 == nil {
        return t2
    }

    if t2 == nil {
        return t1
    }

    t1.Val += t2.Val

    t1.Left = mergeTrees(t1.Left, t2.Left)
    t1.Right = mergeTrees(t1.Right, t2.Right)

    return t1

}