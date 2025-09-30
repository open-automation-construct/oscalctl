package component

import (
	"github.com/spf13/cobra"

	"github.com/open-automation-construct/stigctl/internal/oscal/component"
)

// NewCmd creates a new OSCAL component command
func NewCmd() *cobra.Command {
	var inputFile string
	var outputFile string
	var cciFile string

	cmd := &cobra.Command{
		Use:   "component",
		Short: "Generate an OSCAL component from a STIG checklist",
		Long: `Generate an OSCAL component definition from a STIG checklist.
This command takes a STIG checklist as input and converts it to an OSCAL component definition.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return component.GenerateComponent(inputFile, outputFile, cciFile)
		},
	}

	cmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the STIG checklist file")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "component.json", "Path to the output OSCAL component file")
	cmd.Flags().StringVarP(&cciFile, "cci-file", "cci", "", "Path to the CCI file (optional, will use embedded if not provided)")
	_ = cmd.MarkFlagRequired("input")

	return cmd
}