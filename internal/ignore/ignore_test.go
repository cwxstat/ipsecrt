package ignore

import (
	"testing"
)

func TestGithub(t *testing.T) {
	r, err := Github()
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func Test_githubIPs(t *testing.T) {
	r := githubIPs()

	t.Log(r)
}
