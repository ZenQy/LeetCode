package solve0058

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	assert.Equal(t, lengthOfLastWord("Hello World"), 5)
	assert.Equal(t, lengthOfLastWord("   fly me   to   the moon  "), 4)
	assert.Equal(t, lengthOfLastWord("luffy is still joyboy"), 6)
}
