package medium

import (
	"testing"
)

func Test_Addnumbers(t *testing.T) {
	firstNum := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 2,
		},
	}
	secondNum := ListNode{
		Val: 1,
	}
	resultNum := AddTwoNumbers(&firstNum, &secondNum)
	if resultNum.Val == 3 && resultNum.Next.Val == 2 {
		t.Log("Success")
	} else {
		t.Log("Wrong answer, got: ", resultNum.Val, resultNum.Next.Val)
		t.Fail()
	}
}
