/***********************************************************************************************************
*执行用时：8 ms, 在所有 Go 提交中击败了69.42%的用户
*内存消耗：7 MB, 在所有 Go 提交中击败了8.13%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/odd-even-linked-list
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
struct ListNode* oddEvenList(struct ListNode* head) {
    struct ListNode *odd_head = NULL;
    struct ListNode *even_head = NULL;
    struct ListNode *p1 = NULL;
    struct ListNode *p2 = NULL;
        
    if(head == NULL || head->next == NULL)
        return head;
    p1 = odd_head = head;
    p2 = even_head = head->next;
    
    while(p1 != NULL && p2 != NULL){
        if(p2->next != NULL){
            p1->next = p2->next;
            p1 = p1->next;
        }
        else{
            p1->next = NULL;
            break;
        }
        
        if(p1->next != NULL){
            p2->next = p1->next;
            p2 = p2->next;
        }
        else{
            p2->next = NULL;
            break;
        }
    }
    
    p1->next = even_head;
    return head;
    
}