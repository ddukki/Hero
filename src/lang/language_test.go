package lang

import (
	"fmt"
	"testing"

	"github.com/ddukki/Hero/src/util"
	"github.com/stretchr/testify/assert"
)

func TestNameGen(t *testing.T) {
	assert.Equal(t, 2, len(DB.Languages))
	l := DB.Languages[util.GetRandomKey(DB.Languages)]
	r := l.Rules[util.GetRandomKey(l.Rules)]

	randName := r.Generate()
	fmt.Println(randName)
	assert.Less(t, 0, len(randName))
}
