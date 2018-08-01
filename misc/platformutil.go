package misc

/*
⊕ [Go by Example: Environment Variables](https://gobyexample.com/environment-variables)

*/
import (
	"fmt"
	"os"
	"strings"
)

func SetEnvironmentVariable(name string, value string) string {
	os.Setenv(name, value)
	return os.Getenv(name)
}

func GetEnvironmentVariable(name string) string {
	return os.Getenv(name)
}

func AllEnvironmentVars() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
