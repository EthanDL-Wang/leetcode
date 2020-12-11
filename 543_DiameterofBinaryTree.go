/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了96.26%的用户
*内存消耗：4.6 MB, 在所有 Go 提交中击败了89.41%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
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

 var res int = 0

 func getDepth(root *TreeNode) int {
 
	 if root == nil {
		 return 0
	 }
 
	 left := getDepth(root.Left)
	 right := getDepth(root.Right)
 
	//left为左节点的深度
	//right为右节点的深度
	//left+right为以root为中心的最短路径长度
	 if res < left + right {
		 res = left + right
	 }
 
	 //计算root节点的深度
	 max := int(math.Max(float64(left), float64(right)))+1
	 return max
 
 }
 
 
 
 func diameterOfBinaryTree(root *TreeNode) int {
 
	 res = 0
 
	 getDepth(root)
 
	 return res
 
 
 }