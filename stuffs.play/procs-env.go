package main

import (
	"fmt"

	"github.com/samlet/gostuff/misc"
)

func main() {
	result := misc.SetEnvironmentVariable("hello", "world")
	fmt.Println(result)

	misc.AllEnvironmentVars()
}
