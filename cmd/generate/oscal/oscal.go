package oscal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/open-automation-construct/stigctl/pkg/oscal/component"
)

// NewCmd creates a new OSCAL command
func NewCmd() *cobra.Command {
	oscalCmd := &cobra.Command{
		Use:   "oscal",
		Short: "Generate OSCAL artifacts from STIG data",
		Long:  `Generate various OSCAL artifacts from STIG data, such as component definitions.`,
	}

	// Add subcommand
	oscalCmd.AddCommand(newComponentCmd())

	return oscalCmd
}

// newComponentCmd creates a component subcommand
func newComponentCmd() *cobra.Command {
	componentCmd := &cobra.Command{
		Use:   "component",
		Short: "Generate an OSCAL component definition from a STIG checklist",
		Long: `Generate an OSCAL component definition from a STIG checklist.
This creates a machine-readable representation of the STIG compliance data
that can be used with OSCAL-compatible tools.`,
		RunE: generateOSCALComponent,
	}

	// Add flags
	componentCmd.Flags().StringP("input", "i", "", "Path to the STIG checklist (required)")
	componentCmd.Flags().StringP("output", "o", "", "Path to the output OSCAL component definition (required)")
	componentCmd.Flags().String("cci-map", "", "Path to a custom CCI XML document (optional, uses embedded CCI list if not specified)")

	// Bind flags to viper
	if err := viper.BindPFlag("oscal.component.input", componentCmd.Flags().Lookup("input")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding flag: %v\n", err)
	}
	if err := viper.BindPFlag("oscal.component.output", componentCmd.Flags().Lookup("output")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding flag: %v\n", err)
	}
	if err := viper.BindPFlag("oscal.component.cciMap", componentCmd.Flags().Lookup("cci-map")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding flag: %v\n", err)
	}

	// Mark required flags
	if err := componentCmd.MarkFlagRequired("input"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
	if err := componentCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}

	return componentCmd
}

// generateOSCALComponent handles the generate component command
func generateOSCALComponent(cmd *cobra.Command, args []string) error {
	inputPath := viper.GetString("oscal.component.input")
	outputPath := viper.GetString("oscal.component.output")
	cciPath := viper.GetString("oscal.component.cciMap")

	// Verify input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", inputPath)
	}

	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Verify custom CCI file exists if specified
	if cciPath != "" {
		if _, err := os.Stat(cciPath); os.IsNotExist(err) {
			return fmt.Errorf("specified CCI mapping file does not exist: %s", cciPath)
		}
		fmt.Printf("Using custom CCI mapping file: %s\n", cciPath)
	} else {
		fmt.Println("Using embedded CCI mapping file")
	}

	// Generate the OSCAL component
	if err := component.GenerateComponent(inputPath, outputPath, cciPath); err != nil {
		return fmt.Errorf("failed to generate OSCAL component: %w", err)
	}

	fmt.Printf("Successfully generated OSCAL component: %s\n", outputPath)
	return nil
}