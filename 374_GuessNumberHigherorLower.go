/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

 func guessNumber(n int) int {
    left := 1
    right := n 
    var middle int
    for left <= right {
        middle = left + (right - left)/2

        if guess(middle) == 0 {
            break
        } else if guess(middle) > 0 {
            left = middle + 1
        } else if guess(middle) < 0 {
            right = middle - 1
        }
    }

    return middle
} 