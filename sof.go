package sof

import (
	"github.com/x19290/sof/child"
)

func three() int {
	return one() + child.Two()
}

func Three() int {
	return three()
}
