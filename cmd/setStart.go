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

// setStartCmd represents the setStart command
var (
	startDate string

	setStartCmd = &cobra.Command{
		Use:   "setStart",
		Short: "Set the parameter for contract start date",
		Long: `Contract Start Date will be updated to inputted parameter
Any Reports generated will be under this start date
Please write the date in the format "Feb-24"
Example usage:
matrixTool setStart "Jul-23"`,
		Run: func(cmd *cobra.Command, args []string) {
			if startDate == "" && len(args) != 0 {
				startDate = args[0]
				regex := regexp.MustCompile("^[A-Z][a-z]{2}-[0-9]{2}$")
				res := regex.MatchString(startDate)
				if !res {
					fmt.Println("Please specify a start date with the correct format. Run 'matrixTool setStart --help' for more information.")
					return
				}
				fmt.Println("The entered startDate is: " + startDate + "\n")
			} else {
				fmt.Println("Please specify a start date. Run 'matrixTool setStart --help' for more information.")
				return
			}
			dbModify.SetStartDate(startDate)
		},
	}
)

func init() {
	rootCmd.AddCommand(setStartCmd)
	rootCmd.PersistentFlags().StringVarP(&startDate, "startDate", "", "", "[Required] Set contract start date of matrix pricing desired")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setStartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
