/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"github.com/open-automation-construct/stigctl/cmd/generate"
)


var (

cfgFile string

rootCmd = &cobra.Command{
	Use:   "stigctl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default locations: ., $HOME/.stigctl/)")

	generateCmd := generate.NewCmd()
    rootCmd.AddCommand(generateCmd)

    cobra.OnInitialize(func() {
        if err := initializeConfig(rootCmd); err != nil {
            fmt.Println("Error initializing config:", err)
            os.Exit(1)
        }
    })
}

func initializeConfig(cmd *cobra.Command) error {
// 1. Set up Viper to use environment variables.
viper.SetEnvPrefix("STIGCTL")
// Allow for nested keys in environment variables (e.g. `MYAPP_DATABASE_HOST`)
viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
viper.AutomaticEnv()


// 2. Handle the configuration file.
if cfgFile != "" {
	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)
} else {
	// Search for a config file in default locations.
	home, err := os.UserHomeDir()
	// Only panic if we can't get the home directory.
	cobra.CheckErr(err)

	// Search for a config file with the name "config" (without extension).
	viper.AddConfigPath(".")
	viper.AddConfigPath(home + "/.stigctl")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
}

// 3. Read the configuration file.
// If a config file is found, read it in. We use a robust error check
// to ignore "file not found" errors, but panic on any other error.
if err := viper.ReadInConfig(); err != nil {
	// It's okay if the config file doesn't exist.
	var configFileNotFoundError viper.ConfigFileNotFoundError
	if !errors.As(err, &configFileNotFoundError) {
		fmt.Println("No Config Found")
		return err
	}
}

// 4. Bind Cobra flags to Viper.
// This is the magic that makes the flag values available through Viper.
// It binds the full flag set of the command passed in.
err := viper.BindPFlags(cmd.Flags())
if err != nil {
	return err
}

// This is an optional but useful step to debug your config.
fmt.Println("Configuration initialized. Using config file:", viper.ConfigFileUsed())
return nil
}