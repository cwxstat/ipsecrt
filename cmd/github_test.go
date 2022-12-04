package cmd

import (
	"fmt"
	"github.com/cwxstat/ipsecrt/internal/ignore"
	"strings"
	"testing"
)

func TestGithubAdd(t *testing.T) {
	cmds, err := ignore.RouteAdd()
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for _, v := range cmds {
		c := strings.Fields(v)
		//_, err := route.Run(c[0], c[1:]...)
		fmt.Println(c[0], c[1:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		count++
	}
	fmt.Printf("Added %d zoom routes\n", count)
}

func TestGithubDel(t *testing.T) {
	cmds, err := ignore.RouteDelete()
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for _, v := range cmds {
		c := strings.Fields(v)
		//_, err := route.Run(c[0], c[1:]...)
		fmt.Println(c[0], c[1:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		count++
	}
	fmt.Printf("Added %d zoom routes\n", count)
}
