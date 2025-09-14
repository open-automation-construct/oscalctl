package hello

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Hello() {
	fmt.Println("stigctl says Hello")
}

func HelloCmd() *cobra.Command {

	cmd := &cobra.Command {
		Use: "hello",
		Short: "ask stigctl to say hello",
		Long: `Ask a computer to print hello back to you.`,
		Run: func (cmd *cobra.Command, args []string) {
			Hello()
		},
	
	}
	return cmd
}