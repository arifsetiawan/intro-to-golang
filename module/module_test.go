package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetOK(t *testing.T) {
	value := 9
	assert.Equal(t, 9, value, "Form has 9 fields")
}

func Test_GetFailed(t *testing.T) {
	value := 10
	assert.Equal(t, 9, value, "Form has 9 fields")
}
