/***********************************************************************************************************
*执行用时：16 ms, 在所有 Go 提交中击败了63.80%的用户
*内存消耗：9.7 MB, 在所有 Go 提交中击败了20.49%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
请判断一个链表是否为回文链表。
************************************************************************************************************
*/


/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool isPalindrome(struct ListNode* head) {
    
    int len = 0;
    int i;
    struct ListNode *ptr = NULL;
    struct ListNode *p1 = NULL;
    struct ListNode *p2 = NULL;
    struct ListNode *new_head = NULL;
    if(head == NULL || head->next == NULL)
        return true;
    
    ptr = head;
    
    while(ptr != NULL){
        len++;
        ptr = ptr->next;
    }
    
    ptr = head;
    for(i = 0; i < (len+1)/2; i++){
        ptr = ptr->next;
    }
    
    new_head = ptr;
    while(ptr->next != NULL){
        p1 = ptr->next;
        ptr->next = p1->next;
        p1->next = new_head;
        new_head = p1;
    }
    p1 = head;
    p2 = new_head;
    for(i = 0; i < (len)/2; i++){
        if(p1->val != p2->val){
            return false;
        }
        p1 = p1->next;
        p2 = p2->next;
    }
    
    return true;
}
