package oscal

import (
	"github.com/spf13/cobra"

	"github.com/open-automation-construct/stigctl/src/cmd/generate/oscal/component"
)

// NewCmd creates a new OSCAL command
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oscal",
		Short: "Generate OSCAL artifacts",
		Long:  "Generate OSCAL artifacts from STIG checklists and other inputs",
	}

	// Add subcommands
	cmd.AddCommand(component.NewCmd())

	return cmd
}