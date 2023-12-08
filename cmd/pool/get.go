/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package pool

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
	Short: "Get Leostream pool",
	Long: `Get Leostream pool by ID. For example:
		leostream-admin-cli pool get --pool_id 1`,
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("pool_id")
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

		pool_id, _ := cmd.Flags().GetString("pool_id")

		// Convert the pool_id to a string and call the GetPoolAssignment function with it
		pool, err := client.GetPool(pool_id)
		if err != nil {
			fmt.Println(err)
			return
		}
		bytes, _ := json.MarshalIndent(pool, "", "\t")
		fmt.Println(string(bytes))

	},
}

func init() {
	poolCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
