/***********************************************************************************************************
*执行用时：80 ms, 在所有 Go 提交中击败了9.59%的用户
*内存消耗：2.9 MB, 在所有 Go 提交中击败了37.54%的用户
************************************************************************************************************
*/

/***********************************************************************************************************
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。                                             
************************************************************************************************************
*/

func trap(height []int) int {

    count := 0

    for i := 1; i < len(height)-1; i++ {

        leftMax := height[i]
        for j := 0; j < i; j++ {
            if leftMax < height[j] {
                leftMax = height[j]
            }
        }
        

        rightMax := height[i]
        for j := i+1; j < len(height); j++ {
            if rightMax < height[j] {
                rightMax = height[j]
            }
        }

        if leftMax > height[i] && rightMax > height[i] {
            if leftMax > rightMax {
                count += rightMax - height[i] 
            } else {
                count += leftMax - height[i]
            }
        }
        //fmt.Println(i, leftMax, rightMax, height[i], count)
    }


    return count

}


/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了82.89%的用户
*内存消耗：2.8 MB, 在所有 Go 提交中击败了54.92%的用户
************************************************************************************************************
*/

func trap(height []int) int {

    count := 0

    maxIndex := 0
    for i := 0; i < len(height); i++ {
        if height[maxIndex] < height[i] {
            maxIndex = i
        }
    }

    leftIndex := maxIndex

    for leftIndex > 0 {

        tmpIndex := 0
        for i := 0; i < leftIndex; i++ {
            if height[tmpIndex] < height[i] {
                tmpIndex = i
            }
        }

        for i := tmpIndex+1; i < leftIndex; i++ {
            count += height[tmpIndex] - height[i]
            //fmt.Println(i, height[tmpIndex] - height[i])
        }

        leftIndex = tmpIndex
        
    }

    rightIndex := maxIndex

    for rightIndex < len(height)-1 {

        tmpIndex := rightIndex+1
        for i := tmpIndex; i < len(height); i++ {
            if height[tmpIndex] < height[i] {
                tmpIndex = i
            }
        }

        for i := rightIndex+1; i < tmpIndex; i++ {
            count += height[tmpIndex] - height[i]
            //fmt.Println(i, height[tmpIndex] - height[i])
        }

        rightIndex = tmpIndex
        
    }

    return count
}