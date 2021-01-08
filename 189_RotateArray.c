/***********************************************************************************************************
*执行用时：20 ms, 在所有 Go 提交中击败了53.27%的用户
*内存消耗：7.7 MB, 在所有 Go 提交中击败了28.22%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。                                                 
************************************************************************************************************
*/


void rotate(int* nums, int numsSize, int k){
    int *buff = NULL; 
    int moveTimes = 0;
        
    moveTimes = k%numsSize;
    
    if(moveTimes == 0)
        return nums;
    
    
    buff = (int *)malloc(sizeof(int)*numsSize);
    memset(buff, 0, sizeof(int)*numsSize);
    
    memcpy((buff+moveTimes), nums, sizeof(int)*(numsSize-moveTimes));
    memcpy(buff, (nums+numsSize-moveTimes), sizeof(int)*moveTimes);
    memcpy(nums, buff, (sizeof(int)*numsSize));
    return nums;
}