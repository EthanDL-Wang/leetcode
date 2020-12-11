
/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了73.10%的用户
*内存消耗：3 MB, 在所有 Go 提交中击败了23.15%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个二叉树，原地将它展开为一个单链表。
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
 var buff []*TreeNode

 func recursion(root *TreeNode) {
 
	 if root == nil {
		 return 
	 }
 
	 buff = append(buff, root)
	 
	 recursion(root.Left)
	 recursion(root.Right)
 }
 
 
 
 func flatten(root *TreeNode)  {
 
	 buff = make([]*TreeNode, 0)
	 recursion(root)
 
	 for i := 0; i < len(buff); i++ {
		 buff[i].Left = nil
		 if i == len(buff)-1 {
			 buff[i].Right = nil
		 } else {
			 buff[i].Right = buff[i+1]
		 }
		 
	 }
 
 
 
 }