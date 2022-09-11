package DFA

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	matcher := NewDFAMatcher()

	matcher.Build([]string{"日本鬼子", "日本人"})

	match := matcher.Match("小日本娘们")
	fmt.Println(match)
}
