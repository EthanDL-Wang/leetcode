/***********************************************************************************************************
*执行用时：32 ms, 在所有 Go 提交中击败了78.13%的用户
*内存消耗：7.4 MB, 在所有 Go 提交中击败了25.00%的用户
*1.计算墙的宽度，以及砖的总数量，选取其中小的值作为map的大小
*2.缝隙作为key，出现次数作为value放入map
*3.遍历map，找出出现次数最多value
*4.墙高减掉value，即为穿过的最少的砖的数量
************************************************************************************************************
*/

/***********************************************************************************************************
你的面前有一堵矩形的、由多行砖块组成的砖墙。 这些砖块高度相同但是宽度不同。你现在要画一条自顶向下的、穿过最少砖块的垂线。

砖墙由行的列表表示。 每一行都是一个代表从左至右每块砖的宽度的整数列表。

如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你需要找出怎样画才能使这条线穿过的砖块数量最少，并且返回穿过的砖块数量。

你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。
************************************************************************************************************
*/


func leastBricks(wall [][]int) int {

    lenTmp := len(wall[0])

    wallWidth := 0
    for i := 0; i < lenTmp; i++ {
        wallWidth += wall[0][i]
    }

    brickCount := 0
    for _, v := range wall {
        brickCount += len(v)
    }    

    var seamCount int 

    if wallWidth > brickCount {
        seamCount = brickCount
    } else {
        seamCount = wallWidth
    }

    buffMap := make(map[int]int,seamCount)

    for _, v := range wall {
        widthTmp := 0
        for _, w := range v {
            widthTmp += w
            if widthTmp < wallWidth {
                buffMap[widthTmp]++
            }
            
        }
    }

    var maxSeam int
     for _, v := range buffMap {
        if maxSeam < v {
            maxSeam = v
        }
    }

    return len(wall)-maxSeam


}