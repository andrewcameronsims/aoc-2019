package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	modules := []int{1969}
	got := partTwo(modules)
	assert.Equal(t, 966, got)
}
