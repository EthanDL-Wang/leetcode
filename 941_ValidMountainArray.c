/***********************************************************************************************************
*执行用时：48 ms, 在所有 Go 提交中击败了6.49%的用户
*内存消耗：6.9 MB, 在所有 Go 提交中击败了11.73%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

A.length >= 3
在 0 < i < A.length - 1 条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-mountain-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


bool validMountainArray(int* A, int ASize){
    
    if(A == NULL || ASize < 3){
        return false;
    }

    int index;
    int leftFlag = 0;
    int rightFlag = 0;
    int maxIndex = 0;
    int i;

    for(i = 1; i < ASize; i++){
        if(A[maxIndex] < A[i]){
            maxIndex = i;
        }
    }

    if(maxIndex == 0 || maxIndex == ASize-1){
        return false;
    }


    //to left
    index = maxIndex;
    while(index > 0){
        if(A[index] > A[index-1]){
            index--;
        }
        else{
            return false;
        }
    }

    //to right
    index = maxIndex;
    while(index < ASize-1){
        if(A[index] > A[index+1]){
            index++;
        }
        else{
            return false;
        }
    }

    return true;

}