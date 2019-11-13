package main

import (
	"fmt"
	"github.com/vaind/versioning-test/lib"
)

func main() {
	fmt.Println("loaded", lib.Version())
}
