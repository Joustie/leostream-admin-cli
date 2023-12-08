/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package center

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/joustie/go-leostream-admin-cli/cmd"
)

// centerCmd represents the center command
var centerCmd = &cobra.Command{
	Use:   "center",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use list or get")
	},
}

func init() {
	centerCmd.PersistentFlags().String("center_id", "", "Center_id to get center for")
	cmd.RootCmd.AddCommand(centerCmd)
}
