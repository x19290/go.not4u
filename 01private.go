package parent

import (
	"github.com/x19290/go.not4u/child"
)

func private() string {
	return child.Public()
}
