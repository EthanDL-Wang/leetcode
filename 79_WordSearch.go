/***********************************************************************************************************
*执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
*内存消耗：3.5 MB, 在所有 Go 提交中击败了59.84%的用户
*回溯算法
************************************************************************************************************
*/

/***********************************************************************************************************
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
************************************************************************************************************
*/




var direction = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} 

func check(i, j, k int, board [][]byte, visit [][]bool, word string) bool {
    visit[i][j] = true
    if k == len(word)-1  {
        return board[i][j] == word[k] 
    } 

    //fmt.Printf("borad[%d][%d]=%c word[%d]=%c\n",i,j,board[i][j], k, word[k])
    if board[i][j] == word[k] {
        for d := 0; d < 4; d++ {
            newI := i + direction[d][0]
            newJ := j + direction[d][1]

            if (newI >= 0 && newI < len(board)) && (newJ >= 0 && newJ < len(board[0])) && visit[newI][newJ] != true {
                if check(newI, newJ, k+1, board, visit, word) == true {
                    return true
                }
            }
        }        
    }

    visit[i][j] = false
    return false

}

func exist(board [][]byte, word string) bool {

    visited := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        visited[i] = make([]bool, len(board[0]))
    }


    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if check(i, j, 0, board, visited, word) == true {
                return true
            }
        }
    }

    return false



}