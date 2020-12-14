/***********************************************************************************************************
*解题思路：
*1.遍历每一列，当前列成员所在行由1变为0，则此行为弱行，添加进buff。
*2.第一列和最后一列需要特殊处理
*3.第一列的成员如果为0，则不会被第一步所排序，所以需要特殊处理
*4.最后一列成员如果为1，也不会被第一步所排序，所以需要特殊处理
************************************************************************************************************
*/

func kWeakestRows(mat [][]int, k int) []int {

    row := len(mat)
    col := len(mat[0])
    buff := make([]int, k)

    var index int

    for i := 0; i < row; i++ {
        if mat[i][0] == 0 {
            if index == k {
                break
            }
            buff[index] = i
            index++
        }
    }


    for j := 0; j < col-1; j++ {
        if index == k {
            break
        }
        for i := 0; i < row; i++ {
            if mat[i][j] == 1 && mat[i][j+1] == 0 {
                if index == k {
                    break
                }                
                buff[index] = i
                index++

            }
        }
         
    }

    if index != k {
        for i := 0; i < row; i++ {
            if mat[i][col-1] == 1 {
                if index == k {
                    break
                }
                buff[index] = i
                index++
            }
        }
    }



    return buff

}