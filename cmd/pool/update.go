/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package pool

import (
	"github.com/spf13/cobra"
	"fmt"
	"encoding/json"
	"os"
	"strconv"
	"github.com/joustie/leostream-client-go"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Leostream pool",
	Long: `Update Leostream pool by ID. For example:
		leostream-admin-cli pool update --pool_id 1 --name "My Pool" --display_name "My Pool" --provision_onoff 1 --provision_threshold 1 --provision_vm_name test1 --provision_vm_id 1 --provision_name_next_value 1 --provision_max 1 --center_id 1 --center_type "vmware"`,
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
		name, _ := cmd.Flags().GetString("name")
		display_name, _ := cmd.Flags().GetString("display_name")
		provision_onoff, _ := cmd.Flags().GetString("provision_onoff")
		provision_threshold, _ := cmd.Flags().GetString("provision_threshold")
		provision_vm_id, _ := cmd.Flags().GetString("provision_vm_id")
		provision_vm_name, _ := cmd.Flags().GetString("provision_vm_name")
		provision_name_next_value, _ := cmd.Flags().GetString("provision_name_next_value")
		provision_max, _ := cmd.Flags().GetString("provision_max")
		center_id, _ := cmd.Flags().GetString("center_id")
		center_type, _ := cmd.Flags().GetString("center_type")

		pool, err := client.GetPool(pool_id)
		if err != nil {
			fmt.Println(err)
			return
		}

		// update the pool

		if cmd.Flags().Changed("desktop_attributes"){
		
			// Loop through the desktop_attributes array and create a pool attribute for each one
			for _, v := range desktop_attributes {
			// Create a new Pool Attribute
				var pool_attribute leostream.PoolAttributes
				pool_attribute.Text_to_match = v
				pool_attribute.Vm_table_field= "name"
				pool_attribute.Condition_type = "ew"
			
				// Add the pool attribute to the pool
				pool.Pool_definition.Attributes = append(pool.Pool_definition.Attributes, pool_attribute)
			}
		}

		if cmd.Flags().Changed("name"){
			pool.Name = name
		}

		if cmd.Flags().Changed("display_name"){
			pool.Display_name = display_name
		}
		
		if cmd.Flags().Changed("provision_onoff"){
			pool.Provision.Provision_on_off, _ = strconv.ParseInt(provision_onoff, 10, 64)
		}

		if cmd.Flags().Changed("provision_threshold"){
			pool.Provision.Provision_threshold, _ = strconv.ParseInt(provision_threshold, 10, 64)
		}

		if cmd.Flags().Changed("provision_vm_id"){
			pool.Provision.Provision_vm_id, _ = strconv.ParseInt(provision_vm_id, 10, 64)
		}

		if cmd.Flags().Changed("provision_vm_name"){
			pool.Provision.Provision_vm_name = provision_vm_name
		}

		if cmd.Flags().Changed("provision_name_next_value"){
			pool.Provision.Provision_vm_name_next_value, _ = strconv.ParseInt(provision_name_next_value, 10, 64)
		}

		if cmd.Flags().Changed("provision_max"){
			pool.Provision.Provision_max, _ = strconv.ParseInt(provision_max, 10, 64)
		}

		if cmd.Flags().Changed("center_id"){
			pool.Provision.Center.ID, _ = strconv.ParseInt(center_id, 10, 64)
		}

		if cmd.Flags().Changed("center_type"){
			pool.Provision.Center.Type  = center_type
		}

		if pool.Provision.Center.ID == 0 {
			pool.Provision.Center = nil
		}

		// Update the pool
		stored, err := client.UpdatePool(pool_id, *pool, nil)
		if err != nil {
			fmt.Println("Error updating Pool: ", err)
			return
		}
		bytes, _ := json.MarshalIndent(stored, "", "\t")
		fmt.Println(string(bytes))

	},
}

func init() {
	updateCmd.Flags().StringArrayVarP(&desktop_attributes, "desktop_attributes", "", []string{}, "Desktop attributes which determine the desktops to be provisioned (can be used multiple times)")
	updateCmd.Flags().String("name", "", "Pool name")
	updateCmd.Flags().String("display_name", "", "Pool display name in the UI")
	updateCmd.Flags().String("provision_threshold", "", "Threshold at which to provision desktops")
	updateCmd.Flags().String("provision_onoff", "", "Desktop provisioning enabled or not")
	updateCmd.Flags().String("provision_vm_id", "", "ID of the desktop image to use for provisioning")
	updateCmd.Flags().String("provision_name_next_value", "1", "Next value to use for naming provisioned desktops")
	updateCmd.Flags().String("provision_vm_name", "1", "Name of the desktop instance to be provisioned")
	updateCmd.Flags().String("provision_max", "1", "Maximum number of provisioned desktops")
	updateCmd.Flags().String("center_id", "", "ID of center to which this pool belongs")
	updateCmd.Flags().String("center_type", "", "Type of center to which this pool belongs (amazon, azure, vmware, etc.")
	
	
	poolCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
