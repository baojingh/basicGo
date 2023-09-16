package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/**
  @Author   : bob
  @Datetime : 2023-06-03 下午 11:23
  @File     : root.go
  @Desc     :
*/

var rootCmd = &cobra.Command{
	Use:   "greet",
	Short: "this is a greet tool with cobra",
	Long:  `this is a long description for mytool.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = fmt.Errorf("unrecognized command")
	},
}

func Execute() {
	_ = rootCmd.Execute()
}
