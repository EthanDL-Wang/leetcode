/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了93.21%的用户
*内存消耗：3.1 MB, 在所有 Go 提交中击败了56.34%的用户
*1.没旋转，第一个即为最小值
*2.旋转后，从两头开始找最小值
************************************************************************************************************
*/

/***********************************************************************************************************
*把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
*例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
************************************************************************************************************
*/

func minArray(numbers []int) int {

    var min int = numbers[0]
    len := len(numbers)


    if numbers[0] >= numbers[len-1] {
        for i, j := 0, len-1; i < len-1 && j > 0; i, j = i+1, j-1{
            if numbers[i] > numbers[i+1] {
                min = numbers[i+1]
                break
            }

            if numbers[j-1] > numbers[j] {
                min = numbers[j]
                break
            }

        }
    } else {
        min = numbers[0]
    }
    return min

}