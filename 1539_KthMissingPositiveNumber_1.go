/**********************************************************************************************************
*1.二分法
***********************************************************************************************************
*/


func findKthPositive(arr []int, k int) int {
   
    len := len(arr)

    var num int
    var count int
    var left, middle, right int
    var flag int

    for count < k {
        num++

        if num < arr[0] || num > arr[len-1] {
            count++  
        } else {
            left = 0
            right = len -1
            middle = 0
            for left <= right {
                middle = left + (right - left)/2
                if num == arr[middle] {
                    flag = 1
                    break    
                } else if num < arr[middle] {
                    right = middle -1
                } else if num > arr[middle] {
                    left = middle + 1
                }
            }
            if flag == 0 {
                count++
            } else {
                flag = 0
            }
        }
        
    }

        

    return num
}