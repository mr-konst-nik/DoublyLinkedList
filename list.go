package list

import (
	"errors"
	"fmt"
)

//Node for doubly linked list
type Node struct {
	Prev *Node
	Data interface{}
	Next *Node
}

//DoublyLinkedList is type of doubly linked list
type DoublyLinkedList struct {
	Len  int
	Tail *Node
	Head *Node
}

//AddHeadNode adds Node to the beginning
func (dl *DoublyLinkedList) AddHeadNode(data interface{}) {
	newNode := &Node{
		Data: data,
	}
	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
	} else {
		newNode.Next = dl.Head
		dl.Head.Prev = newNode
		dl.Head = newNode
	}
	dl.Len++
}

//AddTailNode adds Node to the end
func (dl *DoublyLinkedList) AddTailNode(data interface{}) {
	newNode := &Node{
		Data: data,
	}
	if dl.Tail == nil {
		dl.Head = newNode
		dl.Tail = newNode
	} else {
		dl.Tail.Next = newNode
		newNode.Prev = dl.Tail
		dl.Tail = newNode
	}
	dl.Len++
}

//PrintForward prints Nodes from beginning to the end
func (dl *DoublyLinkedList) PrintForward() error {
	if dl.Head == nil {
		return errors.New("list is empty")
	}
	temp := dl.Head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.Data, temp.Prev, temp.Next)
		temp = temp.Next
	}
	return nil
}

//PrintReward prints Nodes from end to the beginning
func (dl *DoublyLinkedList) PrintReward() error {
	if dl.Head == nil {
		return errors.New("list is empty")
	}
	temp := dl.Tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.Data, temp.Prev, temp.Next)
		temp = temp.Prev
	}
	return nil
}

//GetLen prints total lengt of the list
func (dl *DoublyLinkedList) GetLen() int {
	return dl.Len
}
