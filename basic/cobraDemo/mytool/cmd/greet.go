package cmd

/**
  @Author   : bob
  @Datetime : 2023-06-03 下午 11:23
  @File     : greet.go
  @Desc     :
*/

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string
	msg  string
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "say hello to someone",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s,%s\n", msg, name)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
	greetCmd.Flags().StringVar(&name, "name", "bob", "Name for the people")
	greetCmd.MarkFlagRequired("name")
	greetCmd.Flags().StringVar(&msg, "msg", "hello", "Greeting msg")
}
