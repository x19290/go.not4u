package not4u

import (
	child "github.com/x19290/go.not4u/1nested"
)

func private() string {
	return child.Public()
}
