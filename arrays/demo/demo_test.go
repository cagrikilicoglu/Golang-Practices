package demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateValue(t *testing.T){
	arr :=[3]string{"travis","scott","astroworld"}
	UpdateValue(&arr)
	expected:=[3]string{"travis test","scott test","astroworld test"}
	assert.Equal(t, expected, arr)
}