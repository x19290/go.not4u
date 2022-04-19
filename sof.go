package sof

import (
	"github.com/x19290/sof/child"
)

func four() int {
	return one() + one() + child.Two()
}

func Four() int {
	return four()
}
