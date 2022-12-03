package zoom

import (
	"fmt"
	"strings"
	"testing"
)

func Test_zoomIPs(t *testing.T) {
	s := zoomIPs()
	fmt.Println(s)
}

func Test_routeAdd(t *testing.T) {
	mock := func() []string {
		return []string{"34.160.0.0/16"}
	}

	result, err := routeAdd(mock)
	if err != nil {
		t.Error(err)
	}
	//route add  34.160.0.0/16 192.168.1.1
	if !strings.Contains(result[0], "/sbin/route add 34.160.0.0/16 ") {
		t.Errorf("Expected route add. Got %v", result[0])
	}
	fmt.Println(result)
}

func Test_routeDelete(t *testing.T) {
	mock := func() []string {
		return []string{"34.160.0.0/16"}
	}

	result, err := routeDelete(mock)
	if err != nil {
		t.Error(err)
	}
	//route add  34.160.0.0/16 192.168.1.1
	if !strings.Contains(result[0], "/sbin/route delete 34.160.0.0/16 ") {
		t.Errorf("Expected route add. Got %v", result[0])
	}
	fmt.Println(result)
}
