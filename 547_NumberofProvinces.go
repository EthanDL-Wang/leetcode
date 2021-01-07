/***********************************************************************************************************
*执行用时：28 ms, 在所有 Go 提交中击败了63.18%的用户
*内存消耗：6.6 MB, 在所有 Go 提交中击败了26.82%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-provinces
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。                                                 
************************************************************************************************************
*/




func findCircleNum(isConnected [][]int) int {


    buff := make([]int, int(math.Max(float64(len(isConnected)), float64(len(isConnected[0])))))
    res := 0
    index := 1

    for i := 0; i < len(isConnected); i++ {
        for j := 0; j < len(isConnected[0]); j++ {

            if isConnected[i][j] == 1 {
                if buff[i] == 0 && buff[j] == 0 {
                    buff[i] = index
                    buff[j] = index 
                    index++
                } else if buff[i] == 0 && buff[j] > 0 {
                    buff[i] = buff[j]
                } else if buff[i] > 0 && buff[j] == 0 {
                    buff[j] = buff[i]
                } else if buff[i] > 0 && buff[j] > 0 {
                    if buff[i] != buff[j] {
                        tmp := buff[i]
                        for k := 0; k < len(buff); k++ {
                            if buff[k] == tmp {
                                buff[k] = buff[j]
                            }
                        }
                    }
                }
            }

        }
    }
    //fmt.Println(buff)
    province := make([]int, len(buff)+1)

    for k := 0; k < len(buff); k++ {
        province[buff[k]]++
    }    
    //fmt.Println(province)
    for k := 0; k < len(province); k++ {
        if province[k] > 0 {
            res++
        }
    } 

    return res

}