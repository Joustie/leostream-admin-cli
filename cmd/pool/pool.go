/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package pool

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/joustie/go-leostream-admin-cli/cmd"
)

var desktop_attributes []string

// poolCmd represents the pool command
var poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use list, get or update")
	},
}

func init() {
	poolCmd.PersistentFlags().String("pool_id", "", "Pool_id to get pool for")
	cmd.RootCmd.AddCommand(poolCmd)


}
