/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	"strconv"
)

// setMilsCmd represents the setMils command
var (
	milsString string

	setMilsCmd = &cobra.Command{
		Use:   "setMils",
		Short: "Set the parameter for the broker fee in mils",
		Long: `Broker fee will be updated to inputted parameter
Any Reports generated will include this fee and will be adjusted accordingly
Example usage:
matrixTool setMils "15"`,
		Run: func(cmd *cobra.Command, args []string) {
			// check if args exist
			if milsString == "" && len(args) != 0 {
				milsString = args[0]
				mils, err := strconv.ParseFloat(milsString, 32)
				// if cannot parse from string
				if err != nil {
					fmt.Println("Please enter a valid number. Run matrixTool setMils --help for more information.")
					return
				}
				// update parameters
				dbModify.SetMils(float32(mils))
				fmt.Println("The entered amount of mils is: " + milsString + "\n")
			} else {
				fmt.Println("Please specify a mils. Run 'matrixTool setMils --help' for more information.")
				return
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(setMilsCmd)
	setMilsCmd.PersistentFlags().StringVarP(&milsString, "milsString", "", "", "[Required] Set broker fee of matrix pricing desired")
}
