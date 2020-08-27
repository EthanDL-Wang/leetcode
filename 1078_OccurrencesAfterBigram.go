/***********************************************************************************************************
*1.先将text分割为sliceText
*2.记录sliceText中first的下标到buff1
*3.根据buff1中记录的下标+1，寻找紧挨着first的sencond的下标，记录到buff2
*4.buff2的元素+1，即为答案的下标
************************************************************************************************************
*/

func findOcurrences(text string, first string, second string) []string {

    sliceText := strings.Split(text, " ")
    sliceTextLen := len(sliceText)

    buff1 := make([]int, sliceTextLen)
    buff2 := make([]int, sliceTextLen)
    var index1, index2 int

    //fmt.Println(sliceText)
    for i := 0; i < sliceTextLen-1; i++ {
        if strings.Compare(sliceText[i], first) == 0 {
            buff1[index1] = i
            index1++
        }
    }
    //fmt.Println(buff1)
    for i := 0; i < index1; i++ {
        if buff1[i]+1 < sliceTextLen-1 && strings.Compare(sliceText[buff1[i]+1], second) == 0 {
            buff2[index2] = buff1[i]+1
            index2++
        }
    }
    //fmt.Println(buff2)
    thirdBuff := make([]string, index2)

    index1 = 0
    for i := 0; i < index2; i++ {
        if buff2[i] + 1 < sliceTextLen {
            thirdBuff[index1] = sliceText[buff2[i] + 1]
            index1++
        }
    }





    return thirdBuff

}