package qux

import (
	"fmt"

	"gitlab.com/loir402/corge"
)

func Qux() string {
	return fmt.Sprintf("Qux %s", corge.Corge())
}
