package not4u

import (
	child "github.com/x19290/go.not4u/1child"
)

func private() string {
	return child.Public()
}
