/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package gateway

import (
	"github.com/spf13/cobra"
	"fmt"
	"encoding/json"
	"os"
	"strconv"
	"github.com/joustie/leostream-client-go"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Leostream gateway",
	Long: `Update Leostream gateway by ID. For example:
		leostream-admin-cli gateway update --gateway_id 1 --name "My Gateway" --address test.com --address_private 10.1.1.1 --use_src_ip 1 --load_balancer_id 1`,
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
		name, _ := cmd.Flags().GetString("name")
		address, _ := cmd.Flags().GetString("address")
		address_private, _ := cmd.Flags().GetString("address_private")
		use_src_ip, _ := cmd.Flags().GetString("use_src_ip")
		notes,_ := cmd.Flags().GetString("notes")
		 
		gateway, err := client.GetGateway(gateway_id)
		if err != nil {
			fmt.Println(err)
			return
		}

		// update the pool
	
		if cmd.Flags().Changed("name"){
			gateway.Name = name
		}
		
		if cmd.Flags().Changed("address"){
			gateway.Address = address
		}

		if cmd.Flags().Changed("address_private"){
			gateway.Address_private = address_private
		}

		if cmd.Flags().Changed("use_src_ip"){
			gateway.Use_src_ip, _ = strconv.ParseInt(use_src_ip, 10, 64)
		}

		if cmd.Flags().Changed("notes"){
			gateway.Notes = notes
		}

		// Update the gateway
		stored, err := client.UpdateGateway(gateway_id, *gateway, nil)
		if err != nil {
			fmt.Println("Error updating Gateway: ", err)
			return
		}
		bytes, _ := json.MarshalIndent(stored, "", "\t")
		fmt.Println(string(bytes))

	},
}

func init() {
	updateCmd.Flags().String("name", "", "Pool name")
	updateCmd.Flags().String("address", "", "Public address of the gateway")
	updateCmd.Flags().String("address_private", "", "Private address of the gateway")
	updateCmd.Flags().String("use_src_ip", "", "Filter source IP 0,1 or 2 (0=no, 1=yes, 2=yes)")
	updateCmd.Flags().String("notes", "", "Notes")

	gatewayCmd.AddCommand(updateCmd)

}
