func canBeEqual(target []int, arr []int) bool {

    len := len(target)
    buff := make(map[int]int, len)

    for i := 0; i < len; i++ {
        buff[target[i]]++
    }

    for i := 0; i < len; i++ {
        if buff[arr[i]] == 0 {
            return false
        }
        
        buff[arr[i]]--
    }    

    for _, v := range buff {
        if v != 0 {
            return false
        }
    }

    return true
}