/*
 * @lc app=leetcode id=292 lang=golang
 *
 * [292] Nim Game
 *
 * https://leetcode.com/problems/nim-game/description/
 *
 * algorithms
 * Easy (55.72%)
 * Likes:    386
 * Dislikes: 1179
 * Total Accepted:    187.5K
 * Total Submissions: 336.4K
 * Testcase Example:  '4'
 *
 * You are playing the following Nim Game with your friend: There is a heap of
 * stones on the table, each time one of you take turns to remove 1 to 3
 * stones. The one who removes the last stone will be the winner. You will take
 * the first turn to remove the stones.
 *
 * Both of you are very clever and have optimal strategies for the game. Write
 * a function to determine whether you can win the game given the number of
 * stones in the heap.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: false
 * Explanation: If there are 4 stones in the heap, then you will never win the
 * game;
 * No matter 1, 2, or 3 stones you remove, the last stone will always
 * be
 * removed by your friend.
 */
func canWinNim(n int) bool {
	// if n <= 3 {
	// 	return true
	// }
	// array := make([]bool, n+1)
	// array[0] = true
	// array[1] = true
	// array[2] = true
	// array[3] = true
	// for i := 4; i < n+1; i++ {
	// 	array[i] = !(array[i-1] && array[i-2] && array[i-3])
	// }
	// return array[n]
	return n == 0 || n%4 != 0
}

