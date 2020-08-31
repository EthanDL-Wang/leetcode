/***********************************************************************************************************
*执行用时：4 ms, 在所有 Go 提交中击败了75.24%的用户
*内存消耗：3.3 MB, 在所有 Go 提交中击败了21.84%的用户
*1.将s作为key，t作为value，一一映射
*2.将t作为key，s作为value，一一映射
************************************************************************************************************
*/

func isIsomorphic(s string, t string) bool {
    
    len := len(s)
    
    buffMap1 := make(map[byte]byte, len)
    buffMap2 := make(map[byte]byte, len)


    for i, v := range []byte(s) {
        _, ok := buffMap1[v]
        if ok == true {
            if buffMap1[v] != t[i] {
                return false
            }
        } else {
            buffMap1[v] = t[i]
        }
    }

    
    for i, v := range []byte(t) {
        _, ok := buffMap2[v]
        if ok == true {
            if buffMap2[v] != s[i] {
                return false
            }
        } else {
            buffMap2[v] = s[i]
        }
    }


    return true
}