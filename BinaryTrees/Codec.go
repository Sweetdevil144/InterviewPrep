package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	serialized   string
	deserialized *TreeNode
}

func Constructor() Codec {
	return Codec{
		serialized:   "",
		deserialized: nil,
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	var result []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, "null")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	for len(result) > 0 && result[len(result)-1] == "null" {
		result = result[:len(result)-1]
	}

	this.serialized = "[" + strings.Join(result, ",") + "]"
	return this.serialized
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" || data == "" {
		return nil
	}
	data = strings.Trim(data, "[]")
	if data == "" {
		return nil
	}
	values := strings.Split(data, ",")
	if len(values) == 0 {
		return nil
	}
	rootVal, _ := strconv.Atoi(strings.TrimSpace(values[0]))
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) {
			val := strings.TrimSpace(values[i])
			if val != "null" {
				if leftVal, err := strconv.Atoi(val); err == nil {
					node.Left = &TreeNode{Val: leftVal}
					queue = append(queue, node.Left)
				}
			}
			i++
		}

		if i < len(values) {
			val := strings.TrimSpace(values[i])
			if val != "null" {
				if rightVal, err := strconv.Atoi(val); err == nil {
					node.Right = &TreeNode{Val: rightVal}
					queue = append(queue, node.Right)
				}
			}
			i++
		}
	}

	this.deserialized = root
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
