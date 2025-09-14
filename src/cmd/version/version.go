package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Version() {
	fmt.Println("stigctl version 0.0.1")
}

func VersionCmd() *cobra.Command {

	cmd := &cobra.Command {
		Use: "version",
		Short: "Print the version of stigctl",
		Long: `Find the version you are currently using.`,
		Run: func (cmd *cobra.Command, args []string) {
			Version()
		},
	}
	return cmd
    
}