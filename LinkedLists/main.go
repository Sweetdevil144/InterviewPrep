package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {

	// 1 -> 2 -> 3 -> 4 -> 5 -> 6
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = &ListNode{Val: 5}
	list.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	list.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	printListNode(removeNthFromEnd(list, 3))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return head.Next
	}
	length, temp := 0, head
	for temp != nil {
		length++
		temp = temp.Next
	}
	index, temp := length-n, head
	if index == 0 {
		return head.Next
	}
	for index > 1 {
		temp = temp.Next
		index--
	}

	temp.Next = temp.Next.Next
	return head
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// slow-fast method to get middle element
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil
	// Reversing second list
	var prev *ListNode
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	first := head
	second = prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ptr := head
	var prev *ListNode = nil
	for ptr.Next != nil {
		ref := ptr
		ptr = ptr.Next
		ref.Next = prev
		prev = ref
	}
	ptr.Next = prev
	return ptr
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	deque := []int{}
	left, right := 0, 0

	for right < len(nums) {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[right] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)

		if left > deque[0] {
			deque = deque[1:]
		}

		if (right + 1) >= k {
			res = append(res, nums[deque[0]])
			left += 1
		}
		right += 1
	}

	return res
}
