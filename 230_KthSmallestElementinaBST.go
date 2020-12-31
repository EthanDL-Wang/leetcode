/***********************************************************************************************************
*执行用时：8 ms, 在所有 Go 提交中击败了97.51% 的用户
*内存消耗：6.3 MB, 在所有 Go 提交中击败了74.48%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
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


 func reverse(node *TreeNode, k int, res *int, index *int) {

    if node == nil {
        return
    }

    if *index > k {
        return 
    }

    reverse(node.Left, k, res, index)
    *index += 1
    if *index == k {
        *res = node.Val
        return
    }


    reverse(node.Right, k, res, index)
}


func kthSmallest(root *TreeNode, k int) int {

    res := 0
    index := 0

    reverse(root, k, &res, &index)

    return res



}


