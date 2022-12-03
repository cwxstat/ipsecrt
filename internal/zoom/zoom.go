package zoom

import (
	"fmt"
	"github.com/cwxstat/ipsecrt/internal/route"
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

func RouteAdd() ([]string, error) {
	return routeAdd(ZoomIPs, route.DefaultGW)
}

func routeAdd(f func() []string, fgw func() string) ([]string, error) {
	out := []string{}
	var err error

	ip := f()
	gw := fgw()
	if gw == "" {
		return out, fmt.Errorf("no default gateway")
	}
	for _, v := range ip {
		if v == "" {
			err = fmt.Errorf("empty ip")
			continue
		}
		out = append(out, fmt.Sprintf("/sbin/route add %s %s", v, gw))
	}
	return out, err
}

func RouteDelete() ([]string, error) {
	return routeDelete(ZoomIPs, route.DefaultGW)
}

func routeDelete(f func() []string, fgw func() string) ([]string, error) {
	out := []string{}
	var err error

	ip := f()
	gw := fgw()
	if gw == "" {
		return out, fmt.Errorf("no default gateway")
	}
	for _, v := range ip {
		if v == "" {
			err = fmt.Errorf("empty ip")
			continue
		}
		out = append(out, fmt.Sprintf("/sbin/route delete %s %s", v, gw))
	}
	return out, err
}

func ZoomIPs() []string {
	return zoomIPs()
}

func zoomIPs() []string {
	s := Read("Zoom.txt")
	return strings.Split(s, "\n")

}
