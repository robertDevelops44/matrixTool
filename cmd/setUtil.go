/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	"regexp"
)

// setUtilCmd represents the setUtil command
var (
	util string

	setUtilCmd = &cobra.Command{
		Use:   "setUtil",
		Short: "Set the parameter for utility",
		Long: `Utility entered will be updated to inputted parameter
Any Reports generated will be under this utility
Please enter the utility short code (all-capitalized): "APS" rather than "Potomac Edison"
Example usage:
matrixTool setUtil "APS"`,
		Run: func(cmd *cobra.Command, args []string) {
			if util == "" && len(args) != 0 {
				util = args[0]
				regex := regexp.MustCompile("^[A-Z]+$")
				res := regex.MatchString(util)
				if !res {
					fmt.Println("Please specify a utility code with the correct format. Run 'matrixTool setUtil --help' for more information.")
					return
				}
				if dbModify.GetUtilByCode(util) == "" {
					fmt.Println("Utility does not exist. Run matrixTool showUtils for all possible utilities.")
					return
				}
				fmt.Println("The entered utility code is: " + util + "\n")
			} else {
				fmt.Println("Please specify a utility code. Run 'matrixTool setUtil --help' for more information.")
				return
			}
			dbModify.SetUtil(util)

		},
	}
)

func init() {
	rootCmd.AddCommand(setUtilCmd)
	rootCmd.PersistentFlags().StringVarP(&util, "util", "", "", "[Required] Set utility of matrix pricing desired")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setUtilCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setUtilCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
