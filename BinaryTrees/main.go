package main

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// Example test for Codec serialization and deserialization

	// Helper function to build a simple binary tree:
	//        1
	//       / \
	//      2   3
	//         / \
	//        4   5
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	println("Serialized:", data)
	ans := deser.deserialize(data)
	// Print the root value of the deserialized tree to verify
	if ans != nil {
		println("Deserialized root value:", ans.Val)
		println("Deserialized tree structure:")
		printTree(ans, "", true)
	} else {
		println("Deserialized tree is nil")
	}
}

// Helper function to print tree structure
func printTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	println(prefix + "└── " + strconv.Itoa(node.Val))

	if node.Left != nil {
		printTree(node.Left, prefix+"    ", true)
	}
	if node.Right != nil {
		printTree(node.Right, prefix+"    ", false)
	}
}

func maxPathSum(root *TreeNode) int {
	res := []int{root.Val}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := max(dfs(node.Left), 0)
		rightMax := max(dfs(node.Right), 0)
		res[0] = max(res[0], node.Val+leftMax+rightMax)
		return node.Val + max(leftMax, rightMax)
	}
	dfs(root)
	return res[0]
}

func kthSmallest(root *TreeNode, k int) int {
	cnt, res := k, 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		cnt--
		if cnt == 0 {
			res = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left int64, right int64) bool
	dfs = func(node *TreeNode, left int64, right int64) bool {
		if node == nil {
			return true
		}

		val := int64(node.Val)
		if val <= left || val >= right {
			return false
		}

		return dfs(node.Left, left, val) && dfs(node.Right, val, right)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func goodNodes(root *TreeNode) int {
	const maxValue = math.MinInt
	var dfs func(*TreeNode, int, int) int
	dfs = func(node *TreeNode, max, good int) int {
		if node == nil {
			return good
		}
		if node.Val >= max {
			good++
			max = node.Val
		}
		good = dfs(node.Left, max, good)
		good = dfs(node.Right, max, good)
		return good
	}
	return dfs(root, maxValue, 0)
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		r := -1
		temp := []*TreeNode{}
		for i := range queue {
			node := queue[i]
			if node.Left != nil {
				temp = append(temp, node.Left)
				r = node.Left.Val
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
				r = node.Right.Val
			}
		}
		queue = temp
		if len(temp) > 0 {
			res = append(res, r)
		}
	}
	return res
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if leftToRight {
				levelVals[i] = node.Val
			} else {
				levelVals[levelSize-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelVals)
		leftToRight = !leftToRight
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	res = append(res, []int{root.Val})
	for len(queue) > 0 {
		rep := []*TreeNode{}
		temp := []int{}
		for i := range queue {
			if queue[i].Left != nil {
				rep = append(rep, queue[i].Left)
				temp = append(temp, queue[i].Left.Val)
			}
			if queue[i].Right != nil {
				rep = append(rep, queue[i].Right)
				temp = append(temp, queue[i].Right.Val)
			}
		}
		queue = rep
		if len(temp) > 0 {
			res = append(res, temp)
		}
	}
	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := indexOf(inorder, rootVal)
	if index == -1 {
		return nil
	}
	root := TreeNode{Val: rootVal}
	if index+1 > len(preorder) {
		return nil
	}
	root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return &root
}

func indexOf(arr []int, val int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode, res []int) []int

	dfs = func(node *TreeNode, res []int) []int {
		if node != nil {
			res = dfs(node.Left, res)
			res = append(res, node.Val)
			res = dfs(node.Right, res)
		}
		return res
	}
	return dfs(root, res)
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode, res []int) []int
	dfs = func(node *TreeNode, res []int) []int {
		if node != nil {
			res = append(res, node.Val)
			res = dfs(node.Left, res)
			res = dfs(node.Right, res)
		}
		return res
	}
	return dfs(root, res)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if abs(left-right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		diameter = max(diameter, left+right)
		return 1 + max(left, right)
	}

	dfs(root)
	return diameter
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	return root
}
