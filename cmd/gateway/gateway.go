/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package gateway

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/joustie/go-leostream-admin-cli/cmd"
)

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
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
	gatewayCmd.PersistentFlags().String("gateway_id", "", "Gateway_id to get gateway for")
	cmd.RootCmd.AddCommand(gatewayCmd)
}
