package solve0661

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageSmoother(t *testing.T) {
	assert.Equal(t, imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}),
		[][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}})
	assert.Equal(t, imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}),
		[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}})
}
