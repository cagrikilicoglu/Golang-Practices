package guestList

import (
	
	"testing"
)

func TestGenerate(t *testing.T){
	expected:=Generate()
	actual := [10][2]string{}
	t.Error(expected,actual)
}