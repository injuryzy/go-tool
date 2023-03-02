[dfa 算法](https://github.com/injuryzy/go-tool)
[dfa 算法](https://github.com/injuryzy/go-tool)

![](https://cdn.nlark.com/yuque/0/2022/jpeg/1535149/1662914586770-0a6c8f3d-d6cb-4dd8-979b-e838121f5a1a.jpeg)

1. 创建字典树
1. 对输入的词典进行匹配

创建节点  这里的结点就是上面那幅图
```go
package DFA

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

```

AddChild 字典树中添加字符<br />FindChild 查询字符<br />AddWords  添加单词

```go
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

```
Build 构建字典树<br />Match 匹配铭感词  // true 存在 false 不存在<br />这里的Match（）方法可以抽一个接口出来，这样可以自定义我们别的过滤算法

