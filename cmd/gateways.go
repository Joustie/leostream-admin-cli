/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	
	"github.com/spf13/cobra"
	"github.com/joustie/leostream-client-go"
)

// poolsCmd represents the pools command
var gatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host := os.Getenv("LEOSTREAM_API_HOSTNAME")
		username := os.Getenv("LEOSTREAM_API_USERNAME")
		password := os.Getenv("LEOSTREAM_API_PASSWORD")
		
			// Create a new Leostream client using the configuration values
		client, err := leostream.NewClient(&host, &username, &password)
		if err != nil {
			fmt.Println(err)
			return
		}

		gateways, err := client.GetGateways()
		if err != nil {
			fmt.Println(err)
			return
		}
		
		for _, gateway := range gateways {

			bytes, _ := json.MarshalIndent(gateway, "", "\t")
			fmt.Println(string(bytes))
		}
	},
}

func init() {
	rootCmd.AddCommand(gatewaysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// poolsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// poolsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
