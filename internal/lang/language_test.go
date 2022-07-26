package lang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameGen(t *testing.T) {
	assert.Equal(t, 2, len(DB.List))

	randName := DB.List[0].Rules[0].Generate()
	fmt.Println(randName)
	assert.Less(t, 0, len(randName))
}
