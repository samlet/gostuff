package files

/*
⊕ [Create directory and nested directory path in GoLang | MusingsCafe](http://www.musingscafe.com/create-directory-and-nested-directory-path-in-golang/)
⊕ [Golang IsNotExist, os.Stat Use: File or Directory Exists - Dot Net Perls](https://www.dotnetperls.com/isnotexist-go)

*/
import (
	"fmt"
	"os"
)

/*
If you need to create all directories in a path e.g. /user/abc/xyz then use os.MkDirAll. If “abc” and “xyz” did not exist, they will both be created.
*/
func CreateDirectory(directoryPath string) {
	//choose your permissions well
	pathErr := os.MkdirAll(directoryPath, 0777)

	//check if you need to panic, fallback or report
	if pathErr != nil {
		fmt.Println(pathErr)
	}
}

/*
Check if file or directory exists
*/
func CreateDirectoryIfNotExists(directoryPath string) {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		CreateDirectory(directoryPath)
	}
}
