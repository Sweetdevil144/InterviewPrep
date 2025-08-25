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
	// Create nodes for the first linked list (number 342)
	l1 := &ListNode{Val: 2, Next: nil}
	// l1.Next = &ListNode{Val: 4, Next: nil}
	// l1.Next.Next = &ListNode{Val: 3, Next: nil}
	// l1.Next.Next.Next = &ListNode{Val: 9, Next: nil}

	// Create nodes for the second linked list (number 465)
	l2 := &ListNode{Val: 8, Next: nil}
	// l2.Next = &ListNode{Val: 6, Next: nil}
	// l2.Next.Next = &ListNode{Val: 4, Next: nil}

	// Call addTwoNumbers function (342 + 465 = 807)
	result := addTwoNumbers(l1, l2)

	// Print the first number
	fmt.Print("First number: ")
	printListNode(l1)

	// Print the second number
	fmt.Print("Second number: ")
	printListNode(l2)

	// Print the result
	fmt.Print("Result: ")
	printListNode(result)
}

func findDuplicate(nums []int) int {
    for _, num := range nums {
        idx := abs(num) - 1
        if nums[idx] < 0 {
            return abs(num)
        }
        nums[idx] *= -1
    }
    return -1
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	dummy := &ListNode{}
	current, carry := dummy, 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return dummy.Next
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
