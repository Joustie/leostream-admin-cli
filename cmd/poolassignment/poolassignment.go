/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package poolassignment

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/joustie/go-leostream-admin-cli/cmd"
)

// poolassignmentCmd represents the poolassignment command
var poolassignmentCmd = &cobra.Command{
	Use:   "poolassignment",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use list or update")
	},
}

func init() {
	poolassignmentCmd.PersistentFlags().String("policy_id", "", "PolicyID to get pool assignments for")
	poolassignmentCmd.PersistentFlags().String("poolassignment_id", "", "PoolAssignmentID to get pool assignment for")
	poolassignmentCmd.MarkPersistentFlagRequired("policy_id")

	cmd.RootCmd.AddCommand(poolassignmentCmd)
	
}
