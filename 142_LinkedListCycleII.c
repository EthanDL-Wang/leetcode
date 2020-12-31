/***********************************************************************************************************
*执行用时：8 ms, 在所有 Go 提交中击败了82.69%的用户
*内存消耗：6.4 MB, 在所有 Go 提交中击败了84.14%的用户
*1.
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/







/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *detectCycle(struct ListNode *head) {
    struct ListNode *fp = NULL;
    struct ListNode *sp = NULL;
    
    if(head == NULL || head->next == NULL || head->next->next == NULL)
        return NULL;
    
    fp = sp = head;
    
    while(sp->next != NULL && fp->next != NULL && fp->next->next != NULL){
        fp = fp->next->next;
        sp = sp->next;
        if(fp == sp){
            fp = head;
            while(fp != sp){
                fp = fp->next;
                sp = sp->next;
            }
            return sp;
        }
        
    }
    
    
    return NULL;
    
}