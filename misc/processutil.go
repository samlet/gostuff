package misc

/**
⊕ [Go by Example: Spawning Processes](https://gobyexample.com/spawning-processes)

*/
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func Date() string {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	// fmt.Println("> date")
	fmt.Println(string(dateOut))
	return string(dateOut)
}

func Grep(content string, findString string) string {
	grepCmd := exec.Command("grep", findString)
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	// "hello grep\ngoodbye grep"
	grepIn.Write([]byte(content))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep " + findString)
	fmt.Println(string(grepBytes))
	return string(grepBytes)
}

func ListAllFiles() {
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

func EnsureExists(program string) bool {
	_, lookErr := exec.LookPath(program)
	if lookErr != nil {
		panic(lookErr)
	}
	return true
}

/*
Exec ...
⊕ [Go by Example: Exec'ing Processes](https://gobyexample.com/execing-processes)
args := []string{"ls", "-a", "-l", "-h"}
*/
func Exec(program string, args []string) {
	binary, lookErr := exec.LookPath(program)
	if lookErr != nil {
		panic(lookErr)
	}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
