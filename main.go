package main

import "fmt"

//Node for doubly linked list
type Node struct {
	Prev *Node
	Data string
	Next *Node
}

type doublyLinkedList struct {
	Len  int
	Tail *Node
	Head *Node
}

func initDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (dl *doublyLinkedList) AddFrontNodeDL(data string) {
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

func (dl *doublyLinkedList) AddEndNodeDL(data string) {
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

func (dl *doublyLinkedList) TraverseForward() error {
	if dl.Head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := dl.Head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.Data, temp.Prev, temp.Next)
		temp = temp.Next
	}
	fmt.Println()
	return nil
}

func (dl *doublyLinkedList) TraverseReward() error {
	if dl.Head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := dl.Tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.Data, temp.Prev, temp.Next)
		temp = temp.Prev
	}
	fmt.Println()
	return nil
}

func (dl *doublyLinkedList) GetLen() int {
	return dl.Len
}

func main() {
	doublyList := initDoublyList()
	doublyList.AddFrontNodeDL("C")
	doublyList.AddFrontNodeDL("B")
	doublyList.AddFrontNodeDL("A")
	doublyList.AddFrontNodeDL("0")
	doublyList.AddEndNodeDL("D")
	doublyList.AddEndNodeDL("E")

	err := doublyList.TraverseReward()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Len: %v\n", doublyList.GetLen())
}
