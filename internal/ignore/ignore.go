package ignore

import (
	"encoding/json"
	"fmt"
	"github.com/cwxstat/ipsecrt/internal/constants"
	"github.com/cwxstat/ipsecrt/internal/route"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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

type GithubMeta struct {
	VerifiablePasswordAuthentication bool `json:"verifiable_password_authentication"`
	SSHKeyFingerprints               struct {
		Sha256Rsa     string `json:"SHA256_RSA"`
		Sha256Ecdsa   string `json:"SHA256_ECDSA"`
		Sha256Ed25519 string `json:"SHA256_ED25519"`
	} `json:"ssh_key_fingerprints"`
	SSHKeys    []string `json:"ssh_keys"`
	Hooks      []string `json:"hooks"`
	Web        []string `json:"web"`
	API        []string `json:"api"`
	Git        []string `json:"git"`
	Packages   []string `json:"packages"`
	Pages      []string `json:"pages"`
	Importer   []string `json:"importer"`
	Actions    []string `json:"actions"`
	Dependabot []string `json:"dependabot"`
}

func githubFromFile() (*GithubMeta, error) {
	s := Read("data/githubmeta.json")
	var result GithubMeta
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		return nil, err
	}
	return &result, nil

}

func Github() (*GithubMeta, error) {
	resp, err := http.Get(constants.GITHUB)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github: %s", resp.Status)
	}

	var result GithubMeta
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil && err != io.EOF {
		return nil, err
	}

	return &result, nil
}

func githubIPs() []string {
	github, err := Github()
	if err != nil {
		github, err = githubFromFile()
		if err != nil {
			return []string{}
		}
	}

	var ips []string
	for _, v := range github.Git {
		ips = append(ips, v)
	}
	return ips
}

func RouteAdd() ([]string, error) {
	return route.RouteAdd(githubIPs, route.DefaultGW)
}

func RouteDelete() ([]string, error) {
	return route.RouteDelete(githubIPs, route.DefaultGW)
}
