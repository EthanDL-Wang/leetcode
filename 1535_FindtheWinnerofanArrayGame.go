/***********************************************************************************************************
*执行用时：136 ms, 在所有 Go 提交中击败了27.34%的用户
*内存消耗：12.5 MB, 在所有 Go 提交中击败了5.47%的用户
*1.首先k大于两倍的lenarr时，即返回arr中最大值即可
*2.否则，将挑战者放进一个队列。与winner进行比较
************************************************************************************************************
*/

/***********************************************************************************************************
给你一个由 不同 整数组成的整数数组 arr 和一个整数 k 。

每回合游戏都在数组的前两个元素（即 arr[0] 和 arr[1] ）之间进行。比较 arr[0] 与 arr[1] 的大小，较大的整数将会取得这一回合的胜利并保留在位置 0 ，较小的整数移至数组的末尾。
当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家 。

返回赢得比赛的整数。

题目数据 保证 游戏存在赢家。
************************************************************************************************************
*/

type node struct {
    value int
    next *node
}


func getWinner(arr []int, k int) int {

    var winner int = arr[0]    
    lenArr := len(arr)

    if k > 2*lenArr {
        max := arr[0]
        for i := 1; i < lenArr; i++ {
            if max < arr[i] {
                max = arr[i]
            }
        }
        return max
    }   


    head := new(node)
    end := head

    head.value = arr[1]
    head.next = nil

    for i := 2; i < lenArr; i++ {
        tmp := new(node)
        tmp.value = arr[i]
        tmp.next = nil
        end.next = tmp
        end = end.next
    }


    var winCount int 

    for {
        if winner > head.value {
            winCount++

        } else {
            winCount = 1

            tmp := winner
            winner = head.value
            head.value = tmp
        }
        end.next = head
        head = head.next
        end = end.next
        end.next = nil

        if winCount == k {
            break
        }


    }

    return winner
}