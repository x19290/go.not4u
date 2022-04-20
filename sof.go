package sof

import (
	"github.com/x19290/sof/child"
)

func one() int {
	return child.One()
}

func One() int {
	return one()
}
