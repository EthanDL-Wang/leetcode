/***********************************************************************************************************
给定由若干 0 和 1 组成的数组 A。我们定义 N_i：从 A[0] 到 A[i] 的第 i 个子数组被解释为一个二进制数（从最高有效位到最低有效位）。

返回布尔值列表 answer，只有当 N_i 可以被 5 整除时，答案 answer[i] 为 true，否则为 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-prefix-divisible-by-5
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。                                               
************************************************************************************************************
*/


/***********************************************************************************************************
*执行用时：28 ms, 在所有 Go 提交中击败了92.13%的用户
*内存消耗：10.4 MB, 在所有 Go 提交中击败了14.61%的用户
************************************************************************************************************
*/

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
bool* prefixesDivBy5(int* A, int ASize, int* returnSize){
    int tmp = 0;
    int i;
    bool *buff = (int *)malloc(sizeof(bool)*ASize);
    memset(buff, false, sizeof(bool)*ASize);
    
    for(i = 0; i < ASize; i++){
        tmp = ((tmp<<1)|A[i])%5;
        if(tmp == 0){
            buff[i] = true;
        }
    }
    
      
    *returnSize = ASize;
    return buff;
}