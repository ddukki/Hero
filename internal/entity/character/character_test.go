package character

import (
	"fmt"
	"testing"
)

func TestNewRandomCharacter(t *testing.T) {
	c := NewRandomCharacter()
	fmt.Print(c.String())
}
