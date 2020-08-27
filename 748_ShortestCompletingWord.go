/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：4.1 MB, 在所有 Go 提交中击败了100.00%的用户
*解题思路：
*1.桶排序法将license内的每一个有效字母的格式记录到buffLicensePlate，并记录总个数licenseValidLen
*2.words中元素长度 < licenseValidLen，则忽略
*3.若已找到words中符合要求的元素，且当前要检查的words中的元素长度 >= 符合要求的元素，则忽略
*4.将当前要检查的words中的元素的每一个字母的个数使用桶排序的方法放入tmpbuff（使用减法的方式）
*5.如果第四条内的tmpbuff中元素有正数，则说明当前检查的words中的元素不符合要求。记录符合要求的元素的下标
*6.返回下标对应的words中的元素
************************************************************************************************************
*/

func shortestCompletingWord(licensePlate string, words []string) string {

    const LEN int = 26

    buffLicensePlate := make([]int,LEN)
    var licenseValidLen int

    for _, v := range licensePlate {
        if v >= 'a' && v <= 'z' {
            buffLicensePlate[v - 'a']++
            licenseValidLen++
        } else if v >= 'A' && v <= 'Z' {
            buffLicensePlate[v - 'A']++
            licenseValidLen++
        }    
    } 
    //fmt.Println(buffLicensePlate)
    //fmt.Println(licenseValidLen)

    tmpbuff := make([]int,LEN)
    var j int
    var index int = -1

    for i, v := range words {
        copy(tmpbuff,buffLicensePlate)
        //fmt.Println(tmpbuff)
        if len(v) < licenseValidLen{
            continue
        }

        if index >= 0 && len(words[index]) <= len(v) {
            continue
        }

        for _, l := range v {
            tmpbuff[l - 'a']--
        }
        //fmt.Println(tmpbuff)
        for j = 0; j < LEN; j++ {
            if tmpbuff[j] > 0 {
                break
            }
        }

        if j == LEN {
            index = i
        }
        //fmt.Println(index, j)
    }


    return words[index]





















}