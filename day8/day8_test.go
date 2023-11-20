package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage_New(t *testing.T) {
	testCases := []struct {
		desc           string
		encoded        []int
		height         int
		width          int
		expectedLayers [][][]int
	}{
		{
			desc:    "",
			encoded: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
			expectedLayers: [][][]int{
				{
					{1, 2, 3},
					{4, 5, 6},
				},
				{
					{7, 8, 9},
					{0, 1, 2},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			img := NewImage(tC.encoded, 2, 3)
			assert.Equal(t, tC.expectedLayers, img.Layers)
		})
	}
}
