/***********************************************************************************************************
*执行用时：24 ms, 在所有 Go 提交中击败了87.36%的用户
*内存消耗：6.6 MB, 在所有 Go 提交中击败了17.13%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-array-by-parity-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/

//golang
func sortArrayByParityII(A []int) []int {

    answer := make([]int, len(A))

    oddIndex := 1
    evenIndex := 0

    for _, v := range A {
        if v % 2 == 0 {
            answer[evenIndex] = v
            evenIndex += 2
        } else {
            answer[oddIndex] = v
            oddIndex += 2
        }
    }

    return answer

}


//C
/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
 int* sortArrayByParityII(int* A, int ASize, int* returnSize) {
    int *B = NULL;
    int i;
    int oddIndex = 1;
    int evenIndex = 0;
    
    B = (int *)malloc(sizeof(int)*ASize);
    memset(B, 0, sizeof(int)*ASize);
    
    for(i = 0; i < ASize; i++){
        if(A[i]%2 == 1){
            B[oddIndex] = A[i];
            oddIndex = oddIndex+2;
        }
        else{
            B[evenIndex] = A[i];
            evenIndex = evenIndex+2;            
        }        
    }
    *returnSize = ASize;
    return B;
    
}