package route

import (
	"fmt"
	"testing"
)

func MockRun(name string, arg ...string) ([]byte, error) {
	return []byte(`Routing tables

Internet:
Destination        Gateway            Flags           Netif Expire
default            link#18            UCSg             ppp0
default            192.168.1.1        UGScIg            en0
8.8.8.8            link#18            UHW3Ig           ppp0   3528
10.10.1.25         link#18            UHWIig           ppp0
10.10.1.50         link#18            UHWIig           ppp0
17.248.190.207     link#18            UHW3Ig           ppp0   3561
17.253.4.253       link#18            UHW3Ig           ppp0   3500
17.253.16.125      link#18            UHW3Ig           ppp0   3500
34.160/16          192.168.1.1        UGSc              en0
52.31.58.3         link#18            UHWIig           ppp0
54.171.149.88      link#18            UHW3Ig           ppp0   3549
65.8.158.91        link#18            UHW3Ig           ppp0   3547
67.23.55.194       192.168.1.1        UGHS              en0
104.20.232.13      link#18            UHWIig           ppp0
104.20.233.13      link#18            UHWIig           ppp0
`), nil
}

func TestRun(t *testing.T) {
	out, err := Run("ls", "-l")
	if err != nil {
		t.Error("Run", err)
	}
	print(out)
}

func TestStat(t *testing.T) {
	out, err := netStat(MockRun, "/usr/sbin/netstat", "-nr", "-f", "inet")
	if err != nil {
		t.Error("Run", err)
	}
	r := parseNetStat(out)
	fmt.Println(r)
}

func TestDefaultGW(t *testing.T) {

	mock := func() ([][]string, error) {
		out, err := netStat(MockRun, "/usr/sbin/netstat", "-nr", "-f", "inet")
		return parseNetStat(out), err
	}

	gw := defaultGW(mock)
	if gw != "192.168.1.1" {
		t.Error("defaultGW", gw)
	}
	fmt.Println(gw)

}
