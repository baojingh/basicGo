package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

/**
  @Author   : bob
  @Datetime : 2023-06-04 上午 8:22
  @File     : version.go
  @Desc     :
*/

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "mytool version info",
	Run: func(cmd *cobra.Command, args []string) {
		name := "git"
		subname := "version"
		args = append([]string{subname}, args...)
		output := exec.Command(name, args...)
		bytes, _ := output.CombinedOutput()
		_, _ = fmt.Fprint(os.Stdout, string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
