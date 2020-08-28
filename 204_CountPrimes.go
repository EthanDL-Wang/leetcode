/***********************************************************************************************************
*执行用时：1404 ms, 在所有 Go 提交中击败了5.08%的用户
*内存消耗：4.9 MB, 在所有 Go 提交中击败了62.72%的用户
*1.申请一个长度为n的数组
*2.遍历n，判断i是否为质数，如果是，将标记i的倍数为非质数
*3.遍历数组，统计质数的数量
************************************************************************************************************
*/

func isPrime(num int) bool {

    tmp := int(math.Sqrt(float64(num)))

    for i := 2; i <= tmp; i++ {
        if num%i == 0 {
            return false
        }
    }

    return true
}

func countPrimes(n int) int {

    buff := make([]bool, n)
    //fmt.Println(buff)

    for i := 2; i < n; i++ {

        if buff[i] == true {
            continue
        }

        if isPrime(i) == true {
            //fmt.Println(i)
            for j := i * 2; j < n; j += i {
                buff[j] = true
            }
        }
    } 

    var count int 

    for i := 2; i < n; i++ {
        if buff[i] == false {
            count++
        }
    }

    //fmt.Println(buff)
    return count
}