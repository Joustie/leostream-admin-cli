/*
Copyright © 2023 Joost Evertse joustie@gmail.com

*/
package poolassignment

import (
	"github.com/spf13/cobra"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"github.com/joustie/leostream-client-go"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Leostream pool assignment",
	Long: `Update Leostream pool assignment by ID. For example:		
		leostream-admin-cli poolassignment update --policy_id 1 --poolassignment_id 1 --pool_id 1 --protocolplan_id 1 --powerplan_id 1 --releaseplan_id 1 --display_mode 9 --offer_filter 0 --offer_quantity 1 --start_if_stopped 0`,
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("poolassignment_id")
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

		//isMemberOf, _ := cmd.Flags().GetString("isMemberOf") 
		policy_id, _ := cmd.Flags().GetString("policy_id")
		poolassignment_id, _ := cmd.Flags().GetString("poolassignment_id")
		pool_id, _ := cmd.Flags().GetString("pool_id")
		protocolplan_id, _ := cmd.Flags().GetString("protocolplan_id")
		powerplan_id, _ := cmd.Flags().GetString("powerplan_id")
		releaseplan_id, _ := cmd.Flags().GetString("releaseplan_id")
		display_mode, _ := cmd.Flags().GetString("display_mode")
		offer_filter, _ := cmd.Flags().GetString("offer_filter")
		offer_quantity, _ := cmd.Flags().GetString("offer_quantity")
		start_if_stopped, _ := cmd.Flags().GetString("start_if_stopped")

		//Get the poolassignment
		data, err := client.GetPoolAssignment(policy_id,poolassignment_id)
			if err != nil {
				fmt.Println(err)
				return
			}

		//Update the poolassignment 
		
		if cmd.Flags().Changed("protocolplan_id"){
			data.Plan_protocol_id, _ = strconv.ParseInt(protocolplan_id, 10, 64)
		}
		
		if cmd.Flags().Changed("isMemberOf"){
			var pa_filters leostream.PoolAssignmentFilters
			pa_filters.Join = "1"
			var filters []leostream.Filters
			pa_filters.Filters = filters
			
			// Loop through the ismemberof array and create a filter for each one
			for _, v := range ismemberof {
			// Create a new filter
				var json_filter leostream.Filters
				json_filter.Offer_filter_attribute = "isMemberOf"
				json_filter.Offer_filter_condition = "eq"
				json_filter.Offer_filter_value = v
				pa_filters.Filters = append(pa_filters.Filters, json_filter)
			}

			// Add the filters to the poolassignment
			data.Offer_filter_json = &pa_filters
			if cmd.Flags().Changed("offer_quantity"){ 
				data.Offer_quantity, _ = strconv.ParseInt(offer_quantity, 10, 64)
			}
		}

		if cmd.Flags().Changed("pool_id"){
			data.Pool_id, _ = strconv.ParseInt(pool_id, 10, 64)
		}
		if cmd.Flags().Changed("powerplan_id"){
			data.Plan_power_control_id, _ = strconv.ParseInt(powerplan_id, 10, 64)
		}
		if cmd.Flags().Changed("releaseplan_id"){
			data.Plan_release_id, _ = strconv.ParseInt(releaseplan_id, 10, 64)
		}
		if cmd.Flags().Changed("display_mode"){
			data.Display_mode = display_mode
		}

		if cmd.Flags().Changed("offer_filter"){
			data.Offer_filter = offer_filter
		}

		if cmd.Flags().Changed("start_if_stopped"){
			data.Start_if_stopped, _ = strconv.ParseInt(start_if_stopped, 10, 64)
		}

		// Update the poolassignment
		stored, err := client.UpdatePoolAssignment(poolassignment_id, *data, policy_id ,nil)
		if err != nil {
			fmt.Println("Error updating PoolAssignment: ", err)
			return
		}
		bytes, _ := json.MarshalIndent(stored, "", "\t")
		fmt.Println(string(bytes))
			 
	},
}

func init() {
	updateCmd.Flags().StringArrayVarP(&ismemberof, "isMemberOf", "", []string{}, "SAML attribute or AD group to add to the pool assignment(can be used multiple times)")
	updateCmd.Flags().String("pool_id", "", "Pool to associate with  the pool assignment")
	updateCmd.Flags().String("protocolplan_id", "", "Protocol plan to associate with the pool assignment")
	updateCmd.Flags().String("powerplan_id", "", "Power plan to associate with  the pool assignment")
	updateCmd.Flags().String("releaseplan_id", "", "AD group to associate with the pool assignment")
	updateCmd.Flags().String("display_mode", "9", "Way desktop and pool are displayed to the user")
	/*
	Offer filter options:	
	0: Included for all users
	1: Only included if user's AD record matches the offer_filter_json criteria
	2: Only included if current date and time is within the offer_filter_json time ranges
	*/
	updateCmd.Flags().String("offer_filter", "0", "Filter how desktops are offered to the user")
	/*
	Offer quantity options:	
	The number of desktops to offer to a user at login
	*/
	updateCmd.Flags().String("offer_quantity", "1", "Number of desktops are offered to the user")
	/*
	Start_if_stopped  options:
	0: Do not start stopped or suspended desktops
	1: Start stopped or suspended desktops
	*/
	updateCmd.Flags().String("start_if_stopped", "0", "Start stopped or suspended desktops")

	//updateCmd.MarkFlagRequired("policy_id")
	//updateCmd.MarkFlagRequired("poolassignment_id")
	//updateCmd.MarkFlagsRequiredTogether("policy_id", "poolassignment_id")

	poolassignmentCmd.AddCommand(updateCmd)

}
