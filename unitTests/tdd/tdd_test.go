package tdd

import "testing"

func TestVowelCount(t *testing.T){
	actual := VowelCount("I love you")
	expected := uint(5)
	if actual != expected {
		t.Errorf("Expected %d is not equal to actual %d", expected,actual)
	}
}