package solve0541

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStr(t *testing.T) {
	assert.Equal(t, reverseStr("abcdefg", 2), "bacdfeg")
	assert.Equal(t, reverseStr("abcd", 2), "bacd")
}
