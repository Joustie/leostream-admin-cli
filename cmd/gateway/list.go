/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package gateway

import (
	"github.com/spf13/cobra"
	"encoding/json"
	"fmt"
	"os"
	"github.com/joustie/leostream-client-go"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Leostream gateways",
	Long: `List Leostream gateways by ID. For example:
		leostream-admin-cli gateway list`,
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
	gatewayCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
