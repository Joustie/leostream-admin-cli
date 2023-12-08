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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Leostream gateway",
	Long: `Get Leostream gateway by ID. For example:
		leostream-admin-cli gateway get --gateway_id 1`,
	PreRun: func(cmd *cobra.Command, args []string) {
	cmd.MarkFlagRequired("gateway_id")
	},
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

		gateway_id, _ := cmd.Flags().GetString("gateway_id")

		// Convert the pool_id to a string and call the GetPoolAssignment function with it
		gateway, err := client.GetGateway(gateway_id)
		if err != nil {
			fmt.Println(err)
			return
		}
		bytes, _ := json.MarshalIndent(gateway, "", "\t")
		fmt.Println(string(bytes))

	},
}

func init() {
	gatewayCmd.AddCommand(getCmd)
}
