package route

import (
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