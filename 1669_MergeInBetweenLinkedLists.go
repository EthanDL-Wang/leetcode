/***********************************************************************************************************
*执行用时：156 ms, 在所有 Go 提交中击败了9.68%的用户
*内存消耗：7.6 MB, 在所有 Go 提交中击败了8.06%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。

请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-in-between-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    
    list2End := list2

    for ;list2End.Next != nil; list2End = list2End.Next {
        
    }

    aNode := list1
    bNode := list1

    i := 0

    for ;bNode != nil; bNode = bNode.Next {
        i++
        if i == a {
            aNode = bNode
        }

        if i == b+1 {
            break
        }

    }    

    aNode.Next = list2
    list2End.Next = bNode.Next

    return list1

}