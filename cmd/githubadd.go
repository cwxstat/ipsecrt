/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/cwxstat/ipsecrt/internal/ignore"
	"github.com/cwxstat/ipsecrt/internal/route"
	"strings"

	"github.com/spf13/cobra"
)

// githubaddCmd represents the githubadd command
var githubaddCmd = &cobra.Command{
	Use:   "githubadd",
	Short: "route add for github addresses",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmds, err := ignore.RouteAdd()
		if err != nil {
			fmt.Println(err)
		}
		count := 0
		for _, v := range cmds {
			c := strings.Fields(v)
			out, err := route.Run(c[0], c[1:]...)
			if count < 10 {
				fmt.Println(v)
				fmt.Printf("out: %s\n", out)
			}
			if count == 11 {
				fmt.Println("...only first 10 routes are shown")
			}

			if err != nil {
				fmt.Println(err)
				continue
			}
			count++
		}
		fmt.Printf("Added %d routes\n", count)
	},
}

func init() {
	rootCmd.AddCommand(githubaddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubaddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubaddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
