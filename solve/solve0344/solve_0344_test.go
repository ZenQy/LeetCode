package solve0344

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s1)
	reverseString(s2)
	assert.Equal(t, s1, []byte{'o', 'l', 'l', 'e', 'h'})
	assert.Equal(t, s2, []byte{'h', 'a', 'n', 'n', 'a', 'H'})
}
