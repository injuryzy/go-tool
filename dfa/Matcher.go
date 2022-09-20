package dfa

type Matcher interface {
	//构建铭感词
	Build([]string)
	// 匹配
	Match(text string) bool
}
