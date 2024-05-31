package main

import "fmt"

/*
Given an integer x, return true if x is a palindrome, and false otherwise.


Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

    -231 <= x <= 231 - 1

Follow up: Could you solve it without converting the integer to a string?
*/

func main() {
	target := 10
	fmt.Println(isPalindrome(target))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	target, newNum := x, 0

	// Get the modulo of the last number to attach to new number

	for target > 0 {
		modulo := target % 10
		newNum = (newNum * 10) + modulo
		target /= 10
	}

	return newNum == x
}
