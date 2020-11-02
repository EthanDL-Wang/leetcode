/***********************************************************************************************************
*执行用时：16 ms, 在所有 Go 提交中击败了20.78%的用户
*内存消耗：6.1 MB, 在所有 Go 提交中击败了32.39%的用户
*1.首先k大于两倍的lenarr时，即返回arr中最大值即可
*2.否则，将挑战者放进一个队列。与winner进行比较
************************************************************************************************************
*/

/***********************************************************************************************************
给定两个数组，编写一个函数来计算它们的交集。
************************************************************************************************************
*/




int sortArray(int *Array, int Size){
    int i, j;
    int tmp;
    int count = 0;
    for(i = 0; i < Size-1; i++){
        for(j = i+1; j < Size; j++){
            if(Array[i] > Array[j]){
                tmp = Array[i];
                Array[i] = Array[j];
                Array[j] = tmp;
            }
        }        
    }
    return 0;
}

/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* intersection(int* nums1, int nums1Size, int* nums2, int nums2Size, int* returnSize) {
    
    if (nums1 == NULL || nums2 == NULL || nums1Size <= 0 || nums2Size <= 0) {
        *returnSize = 0;
        return NULL;
    }


    int i, j;
    int count = 0;

    int *buff = (int *)malloc(sizeof(int)*nums1Size);
    
    memset(buff, 0, sizeof(int)*nums1Size);
    sortArray(nums1, nums1Size);
    
    for(i = 0; i < nums1Size-1; i++){
        if(nums1[i] == nums1[i+1]){
            continue;
        }
        
        for(j = 0; j < nums2Size; j++){
            if(nums1[i] == nums2[j]){
                buff[count++] = nums1[i];
                break;
            }
        }        
    }
    
    
    for(j = 0; j < nums2Size; j++){
        if(nums1[nums1Size-1] == nums2[j]){
            buff[count++] = nums1[i];
            break;
        }
    }         
    
    *returnSize = count;
    return buff;
    
    
}