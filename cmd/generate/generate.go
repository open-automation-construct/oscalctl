package generate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"github.com/open-automation-construct/oscalctl/internal/cklb"
	"github.com/open-automation-construct/oscalctl/cmd/generate/oscal"
)

// NewCmd creates a new generate command
func NewCmd() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate or validate STIG checklists",
		Long: `Generate or validate STIG checklists in CKLB format.
		
This command helps you work with STIG checklists by validating existing 
checklists or generating new ones based on templates, or converting them to other formats like OSCAL.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("generate called")
		
			cklbFile := viper.GetString("cklbFile")
			fmt.Printf("Processing checklist file: %s\n", cklbFile)
			
			// Create a new checklist instance
			checklist := &cklb.Checklist{}
			
			// Try to load the checklist file
			err := checklist.LoadFromFile(cklbFile)
			if err != nil {
				return fmt.Errorf("error loading checklist: %v", err)
			}
			
			// Validate the checklist
			isValid, errors := checklist.Validate()
			if !isValid {
				fmt.Println("Checklist validation failed with the following errors:")
				for _, err := range errors {
					fmt.Printf("  - %s\n", err)
				}
				return fmt.Errorf("checklist validation failed")
			}
			
			fmt.Println("Checklist is valid!")
			
			// Display some basic info about the checklist
			fmt.Printf("Title: %s\n", checklist.Data.Title)
			fmt.Printf("ID: %s\n", checklist.Data.ID)
			fmt.Printf("Number of STIGs: %d\n", len(checklist.Data.STIGs))
			
			return nil
		},
	}

	// Add subcommands
	generateCmd.AddCommand(oscal.NewCmd())
	
	// Add flags for the generate command
	generateCmd.Flags().StringP("cklbFile", "f", "", "Path of the checklist file to process")

	if err := viper.BindPFlag("cklbFile", generateCmd.Flags().Lookup("cklbFile")); err != nil {
	// Handle the error - in a function that can't return an error, we can panic or log
	// Since this is part of initialization, panicking is reasonable
	panic(fmt.Sprintf("Failed to bind flag 'cklbFile': %v", err))
}

	return generateCmd
}