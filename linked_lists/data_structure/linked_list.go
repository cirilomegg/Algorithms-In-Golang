package data_structure

import (
	"log"
	"strconv"
	"strings"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (list *LinkedList) Append(data int) {
	element := &Node{
		Data: data,
	}

	if list.Head == nil {
		list.Head = element
		list.Tail = element
		return
	}

	list.Tail.Next = element
	list.Tail = element
}

func (list *LinkedList) Contains(data int) bool {
	node := list.Search(data)
	return node != nil
}

func (list *LinkedList) Search(data int) *Node {
	current := list.Head

	for current != nil {
		if current.Data == data {
			return current
		}

		current = current.Next
	}

	return nil
}

func (list *LinkedList) Remove(data int) {
	if list == nil {
		return
	}

	current := list.Head
	var previous *Node

	for current != nil {
		if current.Data == data {
			if current == list.Head {
				list.Head = current.Next
			} else {
				previous.Next = current.Next
			}

			if current == list.Tail {
				list.Tail = previous
			}

			current = nil
			return
		}

		previous = current
		current = current.Next
	}
}

func (list *LinkedList) Reverse() {
	list.Tail = list.Head
	current := list.Head
	var previous *Node
	var next *Node

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	list.Head = previous
}

func (list *LinkedList) Print() {
	if list == nil || list.Head == nil {
		return
	}

	builder := strings.Builder{}

	current := list.Head

	for current != nil {
		builder.WriteString(strconv.Itoa(current.Data))

		if current.Next != nil {
			builder.WriteString(" -> ")
		}

		current = current.Next
	}

	log.Println(builder.String())
}
