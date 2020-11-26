/***********************************************************************************************************
*执行用时：24 ms, 在所有 Go 提交中击败了43.85%的用户
*内存消耗：8.3 MB, 在所有 Go 提交中击败了77.05%的用户
*1.
************************************************************************************************************
*/

/***********************************************************************************************************
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


type MinStack struct {
    Value int
    MinValue int
    Next *MinStack
}


/** initialize your data structure here. */
func Constructor() MinStack {
      
    var stack MinStack
    stack.Next = nil
    return stack
}


func (this *MinStack) Push(x int)  {

    node := new(MinStack)

    node.Value = x

    if this.Next == nil {
        node.MinValue = x
    } else {
        if this.Next.MinValue > x {
            node.MinValue = x
        } else {
            node.MinValue = this.Next.MinValue
        }

    }

    node.Next = this.Next
    this.Next = node

}


func (this *MinStack) Pop()  {

    this.Next = this.Next.Next
}


func (this *MinStack) Top() int {

    return this.Next.Value

}


func (this *MinStack) GetMin() int {

    return this.Next.MinValue

}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */