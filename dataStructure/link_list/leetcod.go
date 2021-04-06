package link_list

type ListNode struct {
	Val  int
	Next *ListNode
}

//203 删除链表中的元素(所有元素)
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: -1}
	dummyHead.Next = head
	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		} else {
			prev = prev.Next //要else，是因为判断的是prev.next，如果删除了后下一个还是要删除的元素。则链表不能移动
		}
	}
	return dummyHead.Next //不能反回head，因为head肯定是有值的,只是next被断开了
}

//递归实现203
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements2(head.Next, val)
	if head.Val == val {
		return head.Next //就是把自己的下一个也就是上一个步骤得到的结果放回
	} else {
		return head
	}
}

//206 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil { //得判断头结点是否为空
		return head
	}
	var pre, cur, next *ListNode
	pre = nil //指向当前结点的前一个结点
	cur = head
	next = head.Next
	for cur.Next != nil {
		cur.Next = pre

		pre = cur
		cur = next
		next = next.Next
	}
	cur.Next = pre
	return cur
}

//递归实现206
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rev := reverseList2(head.Next)
	head.Next.Next = head //现在这么想，假如从第一个开始，后面的链表已经翻转完成了，即第二个结点的next指向null，现在将其指向第一个结点
	head.Next = nil       //第一个结点的next再指向nil
	return rev
}
