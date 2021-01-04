/***********************************************************************************************************
*执行用时：12 ms, 在所有 Go 提交中击败了34.31%的用户
*内存消耗：5.3 MB, 在所有 Go 提交中击败了71.69%的用户
*1.递归
************************************************************************************************************
*/

/***********************************************************************************************************
斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和
************************************************************************************************************
*/


int fib(int N){
    if(N <= 1){
        return N;
    }
    return fib(N-1)+fib(N-2);
}