/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了51.33%的用户
*内存消耗：5.7 MB, 在所有 Go 提交中击败了73.37%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
 

说明：

你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 

进阶：

你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


#define MAX_LEN 100

typedef struct {

    int stack[MAX_LEN];
    int tmp_stack[MAX_LEN];
    int index;

} MyQueue;

/** Initialize your data structure here. */

MyQueue* myQueueCreate() {

    MyQueue *my_stack = NULL; 

    my_stack = (MyQueue *)malloc(sizeof(MyQueue)*1);
    my_stack->index = -1;


    return my_stack;
}

/** Push element x to the back of queue. */
void myQueuePush(MyQueue* obj, int x) {

    obj->stack[++(obj->index)] = x;
    return ;
}

/** Removes the element from in front of queue and returns that element. */
int myQueuePop(MyQueue* obj) {

    int i ;
    int tmp_index = -1;
    int element;

    for(i = obj->index; i >= 0; i--){
        obj->tmp_stack[++tmp_index] = obj->stack[i];
    }

    element = obj->tmp_stack[tmp_index]; 
    tmp_index--;
    obj->index = -1;
    for(i = tmp_index; i >= 0; i--){
        obj->stack[++(obj->index)] = obj->tmp_stack[i];
    }

    return element;


}

/** Get the front element. */
int myQueuePeek(MyQueue* obj) {

    return obj->stack[0];

}

/** Returns whether the queue is empty. */
bool myQueueEmpty(MyQueue* obj) {

    if(obj->index == -1){
        return true;
    }

    return false;
}

void myQueueFree(MyQueue* obj) {

    free(obj);
    return ;
}

/**
 * Your MyQueue struct will be instantiated and called as such:
 * MyQueue* obj = myQueueCreate();
 * myQueuePush(obj, x);
 
 * int param_2 = myQueuePop(obj);
 
 * int param_3 = myQueuePeek(obj);
 
 * bool param_4 = myQueueEmpty(obj);
 
 * myQueueFree(obj);
*/