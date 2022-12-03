package route

import (
	"fmt"
	"testing"
)

func MockRun(name string, arg ...string) ([]byte, error) {
	return []byte(`Routing tables

Internet:
Destination        Gateway            Flags           Netif Expire
default            192.168.1.1        UGScg             en0
127                127.0.0.1          UCS               lo0
127.0.0.1          127.0.0.1          UH                lo0
169.254            link#12            UCS               en0      !
192.168.1          link#12            UCS               en0      !
192.168.1.1/32     link#12            UCS               en0      !
192.168.1.1        20:c0:47:bd:f2:e   UHLWIir           en0   1190
192.168.1.100      1c:1b:61:a9:b2:51  UHLWIi            en0   1131
192.168.1.152      70:3a:ca:56:41:78  UHLWI             en0    394
192.168.1.166      70:3a:ca:bb:41:df  UHLWI             en0      !
192.168.1.209      f0:b3:ec:1a:2a:2e  UHLWIi            en0      !
192.168.1.210      d4:90:9d:cd:c0:11  UHLWI             en0    794
192.168.1.227      90:9c:4b:ce:5f:62  UHLWI             en0      !
192.168.1.235/32   link#12            UCS               en0      !
192.168.1.235      9e:3e:51:8e:de:fe  UHLWI             lo0
192.168.1.255      ff:ff:ff:ff:ff:ff  UHLWbI            en0      !
224.0.0/4          link#12            UmCS              en0      !
224.0.0.251        1:0:5e:0:0:fb      UHmLWI            en0
239.255.255.250    1:0:5e:7f:ff:fa    UHmLWI            en0
255.255.255.255/32 link#12            UCS               en0      !
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
