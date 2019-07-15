package main

import "fmt"

type node struct {
	value int;
	nextNode *node;
}

func reverseLink(head *node) *node {
	if head == nil || head.nextNode == nil {
		return head
	}
	var prev *node = nil
	var pcur *node = head
	var next *node

	for pcur != nil {
		if pcur.nextNode == nil {
			pcur.nextNode = prev
			break
		}
		next = pcur.nextNode
		pcur.nextNode = prev
		prev = pcur
		pcur = next
	}
	head.nextNode = nil
	return pcur
}

func printNode(head *node) {
	for head != nil {
		fmt.Println(head.value)
		head = head.nextNode
	}
}

func main() {
	node1 := new(node)
	node1.value = 1
	node2 := new(node)
	node2.value = 2
	node3 := new(node)
	node3.value = 4
	node1.nextNode = node2
	node2.nextNode = node3
	printNode(node1)
	tmp  := reverseLink(node1)
	printNode(tmp)
}
