package sof

import (
	"github.com/x19290/sof/child"
)

func two() int {
	return one() * child.Two()
}

func Two() int {
	return two()
}
