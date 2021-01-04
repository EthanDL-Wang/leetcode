/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了84.73%的用户
*内存消耗：5.7 MB, 在所有 Go 提交中击败了79.74%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
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
struct ListNode* partition(struct ListNode* head, int x) {
    
    if(head == NULL || head->next == NULL){
        return head;
    }
    int flag = 0;
    struct ListNode *x_ptr = NULL;
    struct ListNode *p1 = NULL;
    struct ListNode *p2 = NULL;
    struct ListNode *p3 = NULL;
    struct ListNode *big_head = (struct ListNode*)malloc(sizeof(struct ListNode));
    struct ListNode *small_head = (struct ListNode*)malloc(sizeof(struct ListNode));
    small_head->next = NULL;
    big_head->next = NULL;


    p1 = big_head;
    p2 = small_head;
    p3 = head;
    
    while(p3 != NULL){
        if(p3->val < x){
            p2->next = p3;
            p3 = p3->next;
            p2 = p2->next;
            p2->next = NULL;
        }
        else{
            p1->next = p3;
            p3 = p3->next;
            p1 = p1->next;
            p1->next = NULL;
        }
        
    }
    
    p2->next = big_head->next;
    return small_head->next;
    
}