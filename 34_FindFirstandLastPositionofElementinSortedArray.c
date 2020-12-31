/***********************************************************************************************************
*执行用时：8 ms, 在所有 Go 提交中击败了91.23%的用户
*内存消耗：7.4 MB, 在所有 Go 提交中击败了30.14%的用户
*1.穷举
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
 int* searchRange(int* nums, int numsSize, int target, int* returnSize){
        
    int *buff = NULL;
    int left = 0;
    int right = numsSize-1;
    int middle;
    
    if((buff = (int *)malloc(sizeof(int)*2)) == NULL){
        perror("malloc");
        return NULL;
    }
    buff[0] = buff[1] = -1;
    *returnSize = 2;
    
    
    if(nums == NULL || numsSize <= 0){
        return buff;
    }
    if(target > nums[right] || target < nums[left]){
        return buff;
    }
    
    if(numsSize == 1 && target == nums[left]){
        buff[0] = buff[1] = 0;
        return buff;
    }
    
    while(left != right){
        middle = left + (right - left)/2;
        printf("%d\n", middle);
        if(target > nums[middle]){
            left = middle+1;
        }
        else if(target < nums[middle]){
            right = middle-1;
        }
        else if(target == nums[middle]){
            buff[0] = buff[1] = middle;
            while((buff[1]+1 <= numsSize-1)&&(nums[buff[1]+1] == target)){
                    buff[1] += 1;
            }
            while((buff[0]-1 >= 0)&&(nums[buff[0]-1] == target)){
                    buff[0] -= 1;
            }
            return buff;
        }
    }
    
    if(target == nums[left]){
        buff[0] = buff[1] = left;
        
    }
    return buff;
    
}