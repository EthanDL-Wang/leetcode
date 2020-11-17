/***********************************************************************************************************
*执行用时：28 ms, 在所有 Go 提交中击败了47.54%的用户
*内存消耗：7.2 MB, 在所有 Go 提交中击败了50.00%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。

另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。

返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/matrix-cells-in-distance-order
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。                                               
************************************************************************************************************
*/


func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

    buff := make([][]int, R)

    max := 0
    min := 0

    for i := 0; i < R; i++ {
        buff[i] = make([]int, C)
        for j := 0; j < C; j++ {
            buff[i][j] = int(math.Abs(float64(i-r0)) + math.Abs(float64(j-c0)))

            if buff[i][j] > max {
                max = buff[i][j]
            }

            if buff[i][j] < min {
                min = buff[i][j]
            } 
        }        
    }

    res := make([][]int, R*C)

    index := 0

    for k := min; k <= max; k++ {
        for i := 0; i < R; i++ {
            for j := 0; j < C; j++ {
                if buff[i][j] == k {
                    res[index] = make([]int, 2)
                    res[index][0] = i
                    res[index][1] = j
                    index++
                }
            }        
        }

    }

    return res

}