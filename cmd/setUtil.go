/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	"strings"
)

// setUtilCmd represents the setUtil command
var (
	util string

	setUtilCmd = &cobra.Command{
		Use:   "setUtil",
		Short: "Set the parameter for utility",
		Long: `Utility entered will be updated to inputted parameter
Any Reports generated will be under this utility
Please enter the utility shorthand code: "APS" rather than "Potomac Edison"
Run matrixTool showUtils for all possible utilities.
Example usage:
matrixTool setUtil "APS"`,
		Run: func(cmd *cobra.Command, args []string) {
			// check if args exist
			if util == "" && len(args) != 0 {
				util = strings.ToUpper(args[0])
				// check if util code is valid
				if dbModify.GetUtilByCode(util) == "" {
					fmt.Println("Utility does not exist. Run matrixTool showUtils for all possible utilities.")
					return
				}
				fmt.Println("The entered utility code is: " + util + "\n")
			} else {
				fmt.Println("Please specify a utility code. Run 'matrixTool setUtil --help' for more information.")
				return
			}
			// update parameters
			dbModify.SetUtil(util)
		},
	}
)

func init() {
	rootCmd.AddCommand(setUtilCmd)
	rootCmd.PersistentFlags().StringVarP(&util, "util", "", "", "[Required] Set utility of matrix pricing desired")
}
