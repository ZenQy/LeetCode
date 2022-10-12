package solve0125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, isPalindrome("A man, a plan, a canal: Panama"), true)
	assert.Equal(t, isPalindrome("race a car"), false)
	assert.Equal(t, isPalindrome(" "), true)
	assert.Equal(t, isPalindrome("3P"), false)
}
