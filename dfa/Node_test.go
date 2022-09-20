package dfa

import (
	"fmt"
	"testing"
)

func TestNode_AddChild(t *testing.T) {
	node := Node{}
	node.AddChild('c')
	node.AddChild('d')
	fmt.Println(node)
}

func TestNode_AddWords(t *testing.T) {
	node := Node{}
	node.AddChild('c')
	child := node.FindChild('c')
	fmt.Println(child)
}

func TestNode_FindChild(t *testing.T) {
	node := &Node{
		End:  false,
		Next: nil,
	}
	node.AddWords("strigrn")
	fmt.Println(node)

}
