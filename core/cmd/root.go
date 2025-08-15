/*
Copyright © 2025 faissal maulana
*/
package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

const VERSION = "0.1.2"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "falcon",
	Short: "Falcon is a CLI GO REST APIs Generator",
	Long: `
Falcon is a CLI tool to generate a blueprint Go REST APIs
following Handler-Service-Repository Pattern.
With Falcon you don't need to start from sratch to
structure your REST APIs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := fang.Execute(context.Background(), rootCmd, fang.WithVersion(VERSION)); err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.falcon.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
