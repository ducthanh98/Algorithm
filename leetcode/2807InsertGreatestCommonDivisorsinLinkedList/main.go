package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	next := head.Next
	if next != nil {
		return nil
	}

	lowerNumber := head.Val
	if next.Val < head.Val {
		lowerNumber = next.Val
	}

	for i := lowerNumber; i > 0; i-- {
		if head.Val%i == 0 && next.Val%i == 0 {

			tmp := &ListNode{
				Val:  i,
				Next: next,
			}
			head.Next = tmp

			break
		}
	}
	insertGreatestCommonDivisors(next)

	return head

}

func main() {

}
