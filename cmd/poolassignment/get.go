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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		// Convert the poolassignmentID to a string and call the GetPoolAssignment function with it
		poolassignment, err := client.GetPoolAssignment(policy_id, poolassignment_id)
		if err != nil {
			fmt.Println(err)
			return
		}
		bytes, _ := json.MarshalIndent(poolassignment, "", "\t")
		fmt.Println(string(bytes))

	},
}

func init() {

	
	poolassignmentCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
