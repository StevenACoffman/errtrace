package main

import (
	"fmt"

	"braces.dev/errtrace/cmd/errtrace/testdata/toolexec-test/p1"
)

func main() {
	if err := callP1(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func callP1() error {
	return p1.WrapP2() // @trace
}
