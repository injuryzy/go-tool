package dfa

type Node struct {
	//结束
	End bool
	//节点
	Next map[rune]*Node
}

// AddChild  add char
func (n *Node) AddChild(c rune) *Node {
	if n.Next == nil {
		n.Next = make(map[rune]*Node)
	}
	// 这个字符存在 直接返回
	if node, ok := n.Next[c]; ok {
		return node
	} else {
		n.Next[c] = &Node{
			End:  false,
			Next: nil,
		}
	}
	return n.Next[c]
}

// FindChild find char
func (n *Node) FindChild(c rune) *Node {
	if n.Next == nil {
		return nil
	}
	if node, ok := n.Next[c]; ok {
		return node
	}
	return nil
}

// AddWords add words
func (n *Node) AddWords(w string) {
	node := n
	r := []rune(w)
	for i, _ := range r {
		node = node.AddChild(r[i])
	}
	node.End = true
}
