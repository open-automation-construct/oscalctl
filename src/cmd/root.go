package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	"github.com/Eric-Domeier/stigctl/src/cmd/hello"
	"github.com/Eric-Domeier/stigctl/src/cmd/version"
)

var rootCmd = &cobra.Command{
  Use:   "stigctl",
  Short: "Stigctl is a fast cli tool to automate stig checklists",
  Long: `A Fast and Flexible CLI tool that can help automate stig checklists and create oscal documentation.
                A hobby project`,
  Run: func(cmd *cobra.Command, args []string) {

  },
}

func init() {
	commands := []*cobra.Command{

		version.VersionCmd(),
		hello.HelloCmd(),
	}

	rootCmd.AddCommand(commands...)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}