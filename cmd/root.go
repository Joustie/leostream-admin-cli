/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	//"github.com/spf13/viper"
)



// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "leostream-admin-cli",
	Short: "A go based CLI for the Leostream Admin API",
	Long: `A go based CLI for the Leostream Admin API:

This CLI is designed to allow for the automation of the Leostream Admin API.
It is designed to be used in a CI/CD pipeline to allow for the automation of
the creation and deletion of Leostream resources.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-leostream-admin-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
