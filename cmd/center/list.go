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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Leostream centers",
	Long: `List Leostream centers by ID. For example:
		leostream-admin-cli center list`,
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
		
			centers, err := client.GetCenters()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, center := range centers {
				bytes, _ := json.MarshalIndent(center, "", "\t")
				fmt.Println(string(bytes))
			}
		
		},
}

func init() {
	centerCmd.AddCommand(listCmd)
}
