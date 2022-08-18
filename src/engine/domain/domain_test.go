package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollBaseStat(t *testing.T) {
	for i := 0; i < 1024; i++ {
		assert.LessOrEqual(t, 3, RollBaseStat())
		assert.GreaterOrEqual(t, 18, RollBaseStat())
	}
}
