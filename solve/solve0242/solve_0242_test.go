package solve0242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, isAnagram("anagram", "nagaram"), true)
	assert.Equal(t, isAnagram("cat", "tad"), false)
	assert.Equal(t, isAnagram("aa", "bb"), false)
}
