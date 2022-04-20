package sof

import (
	"github.com/x19290/sof/child"
)

func five() int {
	return two() + child.Three()
}

func Five() int {
	return five()
}
