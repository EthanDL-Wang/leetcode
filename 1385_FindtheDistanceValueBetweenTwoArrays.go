/***********************************************************************************************************
*解题思路：
*1.先将arr2排序，当作一个一维的坐标，然后再看arr1中的元素，落在这个一维坐标的位置
*2.如果arr1落在排序后arr2的左边，那么只需相减后的绝对值与d比较即可
*3.落在排序后的arr2右边类似2
*4.落在排序后的arr2中间，那么就用二分法查找紧挨着arr1的两个arr2，计算绝对值后与d比较
************************************************************************************************************
*/



import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

    len1 := len(arr1)
    len2 := len(arr2)
    var count int 

    sort.Ints(arr2)

    

    for i := 0; i < len1; i++ {


        if arr1[i] < arr2[0] {
            if int(math.Abs(float64(arr1[i] - arr2[0]))) > d {
                count++
                continue
            }
        }

        if arr1[i] > arr2[len2-1] {
            if int(math.Abs(float64(arr1[i] - arr2[len2-1]))) > d {
                count++
                continue
            }
        }

        left := 0
        right := len2 -1
        var middle int
        for left < right {
            middle = left + (right - left)/2
            if arr1[i] == arr2[middle] {
                break
            } else if arr1[i] > arr2[middle] {
                left = middle
            } else if arr1[i] < arr2[middle] {
                right = middle
            }

            if middle == left + (right - left)/2 {

                if int(math.Abs(float64(arr1[i] - arr2[left]))) > d && int(math.Abs(float64(arr1[i] - arr2[right]))) > d {
                    count++
                }
                break
            }
        }
    }    

    return count
}