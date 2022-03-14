package foo

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	expected := "Foo"
	actual := Foo()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
	// assert.Equal(t, "Foo", Foo(), "they should be equal")
}