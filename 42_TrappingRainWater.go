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