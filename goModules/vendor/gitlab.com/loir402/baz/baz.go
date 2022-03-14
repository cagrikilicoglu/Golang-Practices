package baz

import (
	"fmt"

	"gitlab.com/loir402/corge"
)

func Baz() string {
	return fmt.Sprintf("Baz %s", corge.Corge())
}
