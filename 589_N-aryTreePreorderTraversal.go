/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了68.19%的用户
*内存消耗：4 MB, 在所有 Go 提交中击败了74.28%的用户
*1.递归
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个 N 叉树，返回其节点值的前序遍历。
************************************************************************************************************
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 var res []int

 func reverse(node *Node) {
 
	 if node == nil {
		 return 
	 }
 
	 res = append(res, node.Val)
	 //fmt.Println(res)
	 for i := 0; i < len(node.Children); i++ {
		 reverse(node.Children[i])
	 }
 
 }
 
 
 func preorder(root *Node) []int {
 
	 res = make([]int, 0)
 
	 reverse(root)
 
	 return res
 
 
 }