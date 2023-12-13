/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package poolassignment


import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/joustie/leostream-client-go"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Leostream pool assignment",
	Long: `Delete Leostream pool assignment by ID. For example:
		leostream-admin-cli poolassignment delete --policy_id 1 --poolassignment_id 1`,
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("poolassignment_id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		

		host := os.Getenv("LEOSTREAM_API_HOSTNAME")
		username := os.Getenv("LEOSTREAM_API_USERNAME")
		password := os.Getenv("LEOSTREAM_API_PASSWORD")

		policy_id, _ := cmd.Flags().GetString("policy_id")
		poolassignment_id, _ := cmd.Flags().GetString("poolassignment_id")

		// Create a new Leostream client using the configuration values
		client, err := leostream.NewClient(&host, &username, &password)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Delete the poolsassignment
		err = client.DeletePoolAssignment(poolassignment_id, policy_id, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Deleted pool assignment with ID: " + poolassignment_id)
	},
}

func init() {
	poolassignmentCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
