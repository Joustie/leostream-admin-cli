/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package center

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
	Short: "Get Leostream center",
	Long: `Get Leostream center by ID. For example:
		leostream-admin-cli center get --center_id 1`,
	PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("center_id")
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

		center_id, _ := cmd.Flags().GetString("center_id")

		center, err := client.GetCenter(center_id)
		if err != nil {
			fmt.Println(err)
			return
		}
		bytes, _ := json.MarshalIndent(center, "", "\t")
		fmt.Println(string(bytes))

	},
}

func init() {
	centerCmd.AddCommand(getCmd)
}
