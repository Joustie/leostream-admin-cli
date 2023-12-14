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
	"github.com/jedib0t/go-pretty/v6/table"
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

		policy_id, _ := cmd.Flags().GetString("policy_id")

		// If output_format flag is set to json, then output the results as json else output as text
		var json_format bool
		if cmd.Flags().Changed("json") {
			json_format, _ = cmd.Flags().GetBool("json")
		}  

		// Convert the poolassignmentID to a string and call the GetPoolAssignment function with it	 
		poolassignments, err := client.GetPoolAssignments(policy_id)
		if err != nil {
			fmt.Println(err)
			return
		}

		if json_format {
			bytes, _ := json.MarshalIndent(poolassignments, "", "\t")
			fmt.Println(string(bytes))
		}
		if !json_format {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Poolassignment ID", "Poolassignment name", "Pool ID"})
			for _, poolassignment := range poolassignments {
				t.AppendRow([]interface{}{poolassignment.ID, poolassignment.Pool_name, poolassignment.Pool_id})
			}
			t.Render()
		}
		
	},
}

func init() {
	
	poolassignmentCmd.AddCommand(listCmd)

}
