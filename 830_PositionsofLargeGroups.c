/***********************************************************************************************************
*执行用时：12 ms, 在所有 Go 提交中击败了72.46%的用户
*内存消耗：6.4 MB, 在所有 Go 提交中击败了92.75%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
在一个由小写字母构成的字符串 s 中，包含由一些连续的相同字符所构成的分组。

例如，在字符串 s = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。

分组可以用区间 [start, end] 表示，其中 start 和 end 分别表示该分组的起始和终止位置的下标。上例中的 "xxxx" 分组用区间表示为 [3,6] 。

我们称所有包含大于或等于三个连续字符的分组为 较大分组 。

找到每一个 较大分组 的区间，按起始位置下标递增顺序排序后，返回结果。

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/positions-of-large-groups
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


int getEndIndex(char *s, int startIndex, int *endIndex){
    int len = strlen(s);
    int i = 0;
    
    if((startIndex+1) >= len)
        return -1;
    
    for(i = startIndex+1; i < len; i++){
        if(s[startIndex] != s[i]){
            break;
        }
    }
    *endIndex = i;
    return 0;
    
}




/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** largeGroupPositions(char * S, int* returnSize, int** returnColumnSizes){
    int len = strlen(S);
    int count = 0;
    int i;
    int startIndex = 0;
    int tmpIndex = 0;
    int ret;
    
    if(len <= 2){
        *returnSize = 0;
        *returnColumnSizes = NULL;
        return NULL;
    }
        
    
    while(i < len){
        ret = getEndIndex(S, startIndex, &tmpIndex);
        if(ret != 0){
            break;
        }
        
        if((tmpIndex-startIndex) >= 3){
            count++;
        }
        startIndex = tmpIndex;
        i = startIndex;
    }
    
    if(count <= 0){
        *returnSize = 0;
        *returnColumnSizes = NULL;
        return NULL;        
    }
   
    int **buff = (int **)malloc(sizeof(int*)*count);
    int *colSize = (int *)malloc(sizeof(int)*count);
    for(i = 0; i < count; i++){
        buff[i] = (int *)malloc(sizeof(int)*2);
        memset(buff[i], 0, sizeof(int)*2);
        colSize[i] = 2;
    }
    #if 1
    i = 0;
    startIndex = 0;
    count = 0;
    while(i < len){
        ret = getEndIndex(S, startIndex, &tmpIndex);
        if(ret != 0){
            break;
        }
        if((tmpIndex-startIndex) >= 3){
            buff[count][0] = startIndex;
            buff[count][1] = tmpIndex-1;
            count++;
        }
        startIndex = tmpIndex;
        i = startIndex;
    }
    #endif
    *returnSize = count;
    *returnColumnSizes = colSize;
    return buff;
    
}