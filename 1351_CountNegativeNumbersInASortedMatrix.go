func countNegatives(grid [][]int) int {

    var count int
    var left, middle, right int
    row := len(grid)
    col := len(grid[0])

    for i := 0; i < row; i++ {
        left = 0

        if middle == 0 || middle >= col-1 {
            right = col-1
        } else {
            if grid[i][middle] >= 0 {
                right = middle+1
            } else if grid[i][middle] < 0{
                right = middle
            }

            
        }
        
        
        for left <= right {
            middle = left + (right - left)/2
            if grid[i][middle] == 0 {
                for middle < col && grid[i][middle] == 0{
                    middle++
                }
                break
            } else if grid[i][middle] > 0 {
                left = middle + 1
            } else if grid[i][middle] < 0{
                right = middle - 1
            }
        }

        if middle < col {
            if grid[i][middle] >= 0 {
                count += col - middle - 1
            } else if grid[i][middle] < 0{
                count += col - middle
            }
        }     
      
    }

    return count

}