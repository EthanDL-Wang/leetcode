/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了11.16%的用户
*内存消耗：2.9 MB, 在所有 Go 提交中击败了21.18%的用户
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

func levelOrder(root *TreeNode) [][]int {


    if root == nil {
        return nil
    }

    buff1 := make([]*TreeNode, 0)
    buff2 := make([]*TreeNode, 0)

    res := make([][]int, 0)
    buff1 = append(buff1, root)


    for len(buff1) > 0 {
        tmpBuff := make([]int, 0)
        for _, v := range buff1 {
            if v != nil {
                tmpBuff = append(tmpBuff, v.Val)
            }
            if v.Left != nil {
                buff2 = append(buff2, v.Left)
            }
            if v.Right != nil {
                buff2 = append(buff2, v.Right)
            }
        }
        fmt.Println(tmpBuff)
        if len(tmpBuff) > 0 {
            res = append(res, tmpBuff)
        }
        
        buff1 = make([]*TreeNode, len(buff2))
        
        copy(buff1, buff2)
        //fmt.Println(buff1, buff2)
        buff2 = make([]*TreeNode, 0)
    }

    return res

}