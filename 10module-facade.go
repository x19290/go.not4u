package anyid

import (
	two "example.com/x19290/go.not4u/2"
)

func three() int {
	return one() + two.Two()
}

func Three() int {
	return three()
}
