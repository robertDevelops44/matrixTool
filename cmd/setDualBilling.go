/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

// setDualBillingCmd represents the setDualBilling command
var (
	dualBilling string
	dualParam   bool

	setDualBillingCmd = &cobra.Command{
		Use:   "setDualBilling",
		Short: "Set the parameter to include dual billing",
		Long: `Dual Billing inclusion will be updated to inputted parameter
Any Reports generated will be under this parameter
Yes/True will include Dual Billing, No/False will generate a report without dual billing
Please write the parameter in the format True/False, T/F, Yes/No, or Y/N
Example usage:
matrixTool setDualBilling "No"
matrixTool setDualBilling "true"
matrixTool setDualBilling "T"`,

		Run: func(cmd *cobra.Command, args []string) {
			// check if args provided
			if dualBilling == "" && len(args) != 0 {
				// set dualBilling parameter appropriately
				dualBilling := strings.ToLower(args[0])
				if dualBilling == "false" || dualBilling == "f" || dualBilling == "no" || dualBilling == "n" {
					dualParam = false
					fmt.Println("The entered dual billing option is: " + dualBilling + " = " + strconv.FormatBool(dualParam) + "\n")
				} else if dualBilling == "true" || dualBilling == "t" || dualBilling == "yes" || dualBilling == "y" {
					dualParam = true
					fmt.Println("The entered dual billing option is: " + dualBilling + " = " + strconv.FormatBool(dualParam) + "\n")
				} else {
					fmt.Println("Please specify a dual billing option with the correct format. Run 'matrixTool setDualBilling --help' for more information.")
					return
				}
			} else {
				fmt.Println("Please specify a dual billing option. Run 'matrixTool setUtil --help' for more information.")
				return
			}
			// update parameters
			dbModify.SetDualBilling(dualParam)
		},
	}
)

func init() {
	rootCmd.AddCommand(setDualBillingCmd)
	rootCmd.PersistentFlags().StringVarP(&dualBilling, "dualBilling", "", "", "[Required] Set dualBilling inclusion of matrix pricing desired")
}
