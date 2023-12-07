/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package poolassignment

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
	Short: "List Leostream pool assignments",
	Long: `List Leostream pool assignments by ID. For example:	
		leostream-admin-cli poolassignment list --policy_id 1
		leostream-admin-cli poolassignment list --policy_id 1 --poolassignment_id 1
		`,
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

		// Grab the first argument and test if it is a integer get the gateway by id
		// if it is not an integer get the gateway by name
		// if there are no arguments get all poolassignmentIDs
		// if there is more than one argument, we will error out
		policy_id, _ := cmd.Flags().GetString("policy_id")
		poolassignment_id, _ := cmd.Flags().GetString("poolassignment_id")

		if cmd.Flags().Changed("policy_id") && !cmd.Flags().Changed("poolassignment_id") {
			 
				poolassignments, err := client.GetPoolAssignments(policy_id)
				if err != nil {
					fmt.Println(err)
					return
				}
				for _, poolassignment := range poolassignments {
					bytes, _ := json.MarshalIndent(poolassignment, "", "\t")
					fmt.Println(string(bytes))
				}
			 
		} else if cmd.Flags().Changed("policy_id") && cmd.Flags().Changed("poolassignment_id"){
			 
				 
				// Convert the poolassignmentID to a string and call the GetPoolAssignment function with it
				poolassignment, err := client.GetPoolAssignment(policy_id, poolassignment_id)
				if err != nil {
					fmt.Println(err)
					return
				}
				bytes, _ := json.MarshalIndent(poolassignment, "", "\t")
				fmt.Println(string(bytes))
			
		}

	},
}

func init() {
	
	poolassignmentCmd.AddCommand(listCmd)

}
