/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package poolassignment

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/joustie/go-leostream-admin-cli/cmd"
)

var ismemberof []string

// poolassignmentCmd represents the poolassignment command
var poolassignmentCmd = &cobra.Command{
	Use:   "poolassignment",
	Short: "Manage pool assignments",
	Long: `Manage pool assignments, the poolassignment command can be used to list, get, delete or create and update pool assignments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use 'leostream-admin-cli poolassignment <command>' e.g. list, get, delete, create, update")
	},
}

func init() {

	poolassignmentCmd.PersistentFlags().BoolP("json", "j", false, "Output in JSON")
	poolassignmentCmd.PersistentFlags().String("policy_id", "", "PolicyID to get pool assignments for")
	poolassignmentCmd.PersistentFlags().String("poolassignment_id", "", "PoolAssignmentID to get pool assignment for")
	poolassignmentCmd.MarkPersistentFlagRequired("policy_id")

	cmd.RootCmd.AddCommand(poolassignmentCmd)
	
}
