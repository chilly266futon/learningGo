package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEstimateValue(t *testing.T) {
	t.Run("Small", func(t *testing.T) {
		assert.Equal(t, "small", EstimateValue(9))
	})

	t.Run("Medium", func(t *testing.T) {
		assert.Equal(t, "big", EstimateValue(1000))
	})
}
