package route

import (
	"fmt"
	"os/exec"
	"strings"
)

func Run(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return out, err
}

func netStat(f func(string, ...string) ([]byte, error), name string, arg ...string) ([]byte, error) {

	result, err := f(name, arg...)
	if err != nil {
		return nil, err
	}
	return result, err

}

func parseNetStat(in []byte) [][]string {

	s := strings.Split(string(in), "\n")
	out := make([][]string, 0)
	validData := false
	for _, v := range s {
		if strings.HasPrefix(v, "Destination") {
			validData = true
			continue

		}
		if validData {
			out = append(out, strings.Fields(v))
		}

	}
	return out
}

func NetStat() ([][]string, error) {
	out, err := netStat(Run, "/usr/sbin/netstat", "-nr", "-f", "inet")
	if err != nil {
		return nil, err
	}
	return parseNetStat(out), nil
}

func DefaultGW() string {
	return defaultGW(NetStat)
}

func defaultGW(netstat func() ([][]string, error)) string {
	out, err := netstat()
	if err != nil {
		return ""
	}
	for _, v := range out {
		if v[0] == "default" && !strings.Contains(v[1], "link") {
			return v[1]
		}
	}
	return ""
}

func RouteAdd(f func() []string, fgw func() string) ([]string, error) {
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

func RouteDelete(f func() []string, fgw func() string) ([]string, error) {
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
