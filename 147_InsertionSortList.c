/***********************************************************************************************************
*执行用时：56 ms, 在所有 Go 提交中击败了17.67%的用户
*内存消耗：6.8 MB, 在所有 Go 提交中击败了5.78%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
对链表进行插入排序。


插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

 

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insertion-sort-list
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
struct ListNode* insertionSortList(struct ListNode* head) {
    
    if(head == NULL || head->next == NULL)
        return head;
    
    struct ListNode *new_head = NULL;
    struct ListNode *insert_ptr = NULL;
    struct ListNode *p1 = NULL;
    struct ListNode *p2 = NULL;
    int flag = 0;
    
    new_head = (struct ListNode*)malloc(sizeof(struct ListNode));
    new_head->next = head;
    insert_ptr = head->next;
    head->next = NULL;
    
    while(insert_ptr != NULL){
        flag = 0;
        p2 = insert_ptr;
        insert_ptr = insert_ptr->next;
        p1 = new_head;
        while(p1->next != NULL){
            if(p1->next->val >= p2->val){
                p2->next = p1->next;
                p1->next = p2;
                flag = 1;
                break;
            }
            p1 = p1->next;
        }
        if(flag == 0){
            p2->next = NULL;
            p1->next = p2;
        }
    }
    return new_head->next;
    
}