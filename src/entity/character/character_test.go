package character

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomCharacter(t *testing.T) {
	c := NewRandomCharacter()
	fmt.Print(c.String())

	assert.NotNil(t, c.baseStats)
	assert.NotNil(t, c.education)
}
