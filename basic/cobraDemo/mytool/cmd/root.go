package cmd

import "github.com/spf13/cobra"

/**
  @Author   : bob
  @Datetime : 2023-06-03 下午 11:23
  @File     : root.go
  @Desc     :
*/

var rootCmd = &cobra.Command{
	Use:   "greet",
	Short: "this is a greet tool with cobra",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
