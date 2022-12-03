package zoom

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the "this" directory.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}

func Read(rel string) string {
	b, err := os.ReadFile(Path(rel))
	if err != nil {
		return ""
	}
	return string(b)
}

func zoomIPs() []string {
	s := Read("Zoom.txt")
	return strings.Split(s, "\n")

}
