package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {

	tempSlice := []int{10, 20, 30, 40}

	doublyList := &DoublyLinkedList{}
	doublyList.AddHeadNode("E")
	doublyList.AddHeadNode("B")
	doublyList.AddHeadNode("A")
	doublyList.AddHeadNode(0)
	doublyList.AddTailNode("D")
	doublyList.AddTailNode(tempSlice)

	assert.Equal(t, doublyList.Head.Data, 0)
	assert.Equal(t, doublyList.Tail.Data, tempSlice)
	assert.Equal(t, doublyList.Head.Data, doublyList.Head.Next.Prev.Data)
	assert.Equal(t, doublyList.Tail.Data, doublyList.Tail.Prev.Next.Data)
	assert.Equal(t, 6, doublyList.GetLen())

	if err := doublyList.PrintForward(); err != nil {
		fmt.Println(err)
	}

	//Output must be in the following sequence
	//value = 0, prev = <nil>, next = &{0xc000028700 A 0xc0000286c0}
	//value = A, prev = &{<nil> 0 0xc0000286e0}, next = &{0xc0000286e0 B 0xc0000286a0}
	//value = B, prev = &{0xc000028700 A 0xc0000286c0}, next = &{0xc0000286c0 E 0xc000028720}
	//value = E, prev = &{0xc0000286e0 B 0xc0000286a0}, next = &{0xc0000286a0 D 0xc000028740}
	//value = D, prev = &{0xc0000286c0 E 0xc000028720}, next = &{0xc000028720 [10 20 30 40] <nil>}
	//value = [10 20 30 40], prev = &{0xc0000286a0 D 0xc000028740}, next = <nil>

}
