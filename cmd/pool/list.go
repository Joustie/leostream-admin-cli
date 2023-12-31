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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Leostream pools",
	Long: `List Leostream pools by ID. For example:
		leostream-admin-cli pool list`,
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
		
			pools, err := client.GetPools()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, pool := range pools {
				bytes, _ := json.MarshalIndent(pool, "", "\t")
				fmt.Println(string(bytes))
			}
		
		},
}

func init() {

	poolCmd.AddCommand(listCmd)

}
