package DFA

type DFAMatcher struct {
	Root *Node
}

func NewDFAMatcher() *DFAMatcher {
	return &DFAMatcher{
		Root: &Node{
			End: false,
		},
	}
}
func (D *DFAMatcher) Build(strings []string) {
	for i := range strings {
		D.Root.AddWords(strings[i])
	}
}

//  Match 匹配
func (D *DFAMatcher) Match(text string) bool {
	runes := []rune(text)
	child := D.Root
	for i := 0; i < len(runes); i++ {
		//如果没有 ，就往下面找
		findChild := child.FindChild(runes[i])
		if findChild == nil {
			//如果没有匹配 在差从根节点查询
			node := D.Root.FindChild(runes[i])
			if node == nil {
				continue
			}
			//把当前节点给查询节点
			child = node
			continue
		}
		if findChild.End == true {
			return true
		}
		//把根节点换成当前节点
		child = findChild
	}
	return false
}
