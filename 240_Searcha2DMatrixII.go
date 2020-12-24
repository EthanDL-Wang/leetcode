/***********************************************************************************************************
*执行用时：32 ms, 在所有 Go 提交中击败了26.32%的用户
*内存消耗：6.6 MB, 在所有 Go 提交中击败了71.20%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/


func searchMatrix(matrix [][]int, target int) bool {
    
    if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
        return false
    }
    
    
    row := len(matrix)-1
    col := len(matrix[0])-1


    for i := 0; i < len(matrix); i++{
        if target == matrix[i][0] {
            return true
        } else if target < matrix[i][0] {
            row = i - 1
            break
        }
    }

    for i := 0; i < len(matrix[0]); i++{
        if target == matrix[0][i] {
            return true
        } else if target < matrix[0][i] {
            col = i - 1
            break
        }
    }

    //fmt.Println(row,col)

    for i := row; i >= 0; i-- {
        for j := col; j >= 0; j-- {
            //fmt.Println(matrix[i][j])
            if target == matrix[i][j] {
                return true
            }
        }
    }


    return false

}