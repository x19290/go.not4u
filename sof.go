package sof

import (
	"github.com/x19290/sof/nested"
)

func three() int {
	return one() + nested.Two()
}

func Three() int {
	return three()
}
