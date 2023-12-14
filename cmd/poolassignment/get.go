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
	"reflect"
	"github.com/jedib0t/go-pretty/v6/table"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Leostream pool assignment",
	Long: `Get Leostream pool assignment by ID. For example:
		leostream-admin-cli poolassignment get --policy_id 1 --poolassignment_id 1`,
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("poolassignment_id")
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

		policy_id, _ := cmd.Flags().GetString("policy_id")
		poolassignment_id, _ := cmd.Flags().GetString("poolassignment_id")

		// If output_format flag is set to json, then output the results as json else output as text
		var json_format bool
		if cmd.Flags().Changed("json") {
			json_format, _ = cmd.Flags().GetBool("json")
		}

		// Convert the poolassignmentID to a string and call the GetPoolAssignment function with it
		poolassignment, err := client.GetPoolAssignment(policy_id, poolassignment_id)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		if json_format {
			bytes, _ := json.MarshalIndent(poolassignment, "", "\t")
			fmt.Println(string(bytes))
		}

		if !json_format {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Key", "Value"})
			values := reflect.ValueOf(*poolassignment)
			types := values.Type()
			for i := 0; i < values.NumField(); i++ {
				t.AppendRow([]interface{}{types.Field(i).Name, values.Field(i) })
			}
			t.Render()
		}
	},
}

func init() {

	poolassignmentCmd.AddCommand(getCmd)

}
