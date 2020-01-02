// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Example:
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

package medium

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode - ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers - AddTwoNumbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	entry := l1
	carry := 0
	for l1 != nil {
		if l2 != nil {
			(*l1).Val += l2.Val + carry
			if l1.Val > 10 {
				carry = 1
				(*l1).Val -= 10
			}
			fmt.Println(l1.Val)
			if l1.Next == nil {
				l1 = l2.Next
			} else {
				l1 = l1.Next
				l2 = l2.Next
			}
		} else {
			(*l1).Val += carry
			l1 = l1.Next
		}
	}
	return entry
}
