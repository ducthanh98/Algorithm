package main



  type ListNode struct {
      Val int
      Next *ListNode
  }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode = nil
	var head *ListNode =  nil
	for list1 != nil {

		if list1.Val < list2.Val && result == nil{
			tmp := &ListNode{Val: list1.Val,Next: nil}
			head = tmp
			result = tmp
		} else if list1.Val < list2.Val || list2 == nil {
			tmp := &ListNode{Val: list1.Val,Next: nil}
			result.Next = tmp
			result = result.Next
			list1 = list1.Next
		} else {
			tmp := &ListNode{Val: list2.Val,Next: nil}
			result.Next = tmp
			result = result.Next
			list2 = list2.Next
		}

	}

	for list2 != nil {

			tmp := &ListNode{Val: list2.Val,Next: nil}
			result.Next = tmp
			result = result.Next
			list2 = list2.Next

	}
	return head
}

func main(){

}