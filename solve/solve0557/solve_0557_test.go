package solve0557

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, reverseWords("Let's take LeetCode contest"), "s'teL ekat edoCteeL tsetnoc")
	assert.Equal(t, reverseWords( "God Ding"), "doG gniD")
	assert.Equal(t, reverseWords( "  God  Ding  "), "  doG  gniD  ")
}

