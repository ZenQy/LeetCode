package solve0049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	assert.Equal(t, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}),
		[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}})
	assert.Equal(t, groupAnagrams([]string{""}),
		[][]string{{""}})
	assert.Equal(t, groupAnagrams([]string{"e"}),
		[][]string{{"e"}})
}
