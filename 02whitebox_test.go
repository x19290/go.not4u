package simplest

import (
	"github.com/stretchr/testify/assert"
	. "github.com/x19290/go.not4u/1testdata"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, Version, private())
}
