/***********************************************************************************************************
*执行用时：44 ms, 在所有 Go 提交中击败了90.44%的用户
*内存消耗：8.3 MB, 在所有 Go 提交中击败了52.43%的用户
*1.
************************************************************************************************************
*/

/***********************************************************************************************************
给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。
换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。
************************************************************************************************************
*/



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var countA, countB int 

    for tmp := headA; tmp != nil; tmp = tmp.Next {
        countA++
    }

    for tmp := headB; tmp != nil; tmp = tmp.Next {
        countB++
    }

    tmpA := headA
    tmpB := headB

    if countA > countB {
        for i := 0; i < countA-countB; i++ {
            tmpA = tmpA.Next
        }
    } else if countA < countB {
        for i := 0; i < countB-countA; i++ {
            tmpB = tmpB.Next
        }
    }

    for tmpA != nil && tmpB != nil {
        if tmpA == tmpB {
            return tmpA
        }

        tmpA = tmpA.Next
        tmpB = tmpB.Next
    }
    
    return nil

}