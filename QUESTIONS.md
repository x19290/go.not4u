# Questions to Stack Overflow

## 2022-04-20

Subject: publishing a nested module -- publishing or importing is wrong

I have a local directory/module called "sof."
sof has a child module called "sof/child."

sof.Three() referes to sof.three() that referes to sof.one() and child.Two().
Also sof.Test() referes to sof.three()

The current `github.com/x19290/sof` holds the copy of sof.

In that directory, I can do `go test .`

And from a local neighbor of sof (../sofmain), I can do `go run .`

But I can not import real github.com/x19290/sof from sofmain.
I removed `replace` and `require` lines from go.mod
and did `go mod tidy`; then I saw a message:

The contents of sof/sofmain are as follows:

-- sof/sof.go --
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

-- sof/private.go --
package sof

func one() int {
	return 1
}

-- sof/mod_test.go --
package sof

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 3, three())
}

-- sof/go.mod --
module github.com/x19290/sof

go 1.18

replace github.com/x19290/sof/child => ./child

require (
	github.com/stretchr/testify v1.7.1
	github.com/x19290/sof/child v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

-- sof/child/child.go --
package child

func Two() int {
	return 2
}

-- sof/child/go.mod --
module github.com/x19290/sof/child

go 1.18

-- sofmain/sofmain.go --
package main

import (
	"fmt"
	"github.com/x19290/sof"
)

func main() {
	fmt.Println(sof.Three())
}

-- sofmain/go.mod --
module sofmain

go 1.18

replace (
	github.com/x19290/sof => ../sof
	github.com/x19290/sof/child => ../sof/child
)

require github.com/x19290/sof v0.0.0-00010101000000-000000000000

require github.com/x19290/sof/child v0.0.0-00010101000000-000000000000 // indirect

