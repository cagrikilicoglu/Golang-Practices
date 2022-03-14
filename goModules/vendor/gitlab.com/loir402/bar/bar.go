package bar

import (
	"fmt"

	"gitlab.com/loir402/baz"
	"gitlab.com/loir402/qux"
)

// Bar returns the string "Bar"
func Bar() string {
	return fmt.Sprintf("Bar %s %s", baz.Baz(), qux.Qux())
}

func Bar2() string {
	return fmt.Sprintf("Bar2 %s %s", baz.Baz(), qux.Qux())
}
