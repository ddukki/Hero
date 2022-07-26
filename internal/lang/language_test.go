package lang

import (
	"fmt"
	"testing"

	"github.com/ddukki/Hero/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestNameGen(t *testing.T) {
	assert.Equal(t, 2, len(DB.List))
	l := DB.List[util.RandIdx(DB.List)]
	r := l.Rules[util.RandIdx(l.Rules)]

	randName := r.Generate()
	fmt.Println(randName)
	assert.Less(t, 0, len(randName))
}
