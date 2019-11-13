package main

import (
	"fmt"
	lib "github.com/vaind/versioning-test/v2"
)

func main() {
	fmt.Println("loaded", lib.Version())
}
