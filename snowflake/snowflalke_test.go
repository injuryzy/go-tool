package snowflake

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	snowflake := Snowflake{}
	val := snowflake.NextVal()

	fmt.Println(val)
}
