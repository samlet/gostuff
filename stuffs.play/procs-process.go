package main

import (
	"fmt"

	"github.com/samlet/gostuff/misc"
)

func main() {
	fmt.Println(misc.Date())
	misc.Grep("hello grep\ngoodbye grep", "hello")

	misc.ListAllFiles()
	misc.EnsureExists("ls")

	args := []string{"ls", "-a", "-l", "-h"}
	misc.Exec("ls", args)
}
