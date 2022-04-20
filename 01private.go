package sof

import (
	"github.com/x19290/sof/child"
)

func private() string {
	return child.Public()
}
