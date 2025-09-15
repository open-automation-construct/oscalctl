/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checklistCmd represents the checklist command
var checklistCmd = &cobra.Command{
	Use:   "checklist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("checklist called")
	
	name := viper.GetString("name")
	fmt.Printf("Validating a file for you: %s\n", name)
	return nil
	},
}

func init() {
	rootCmd.AddCommand(checklistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checklistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checklistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
