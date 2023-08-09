/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if milsString == "" && len(args) != 0 {
				milsString = args[0]
				mils, err := strconv.ParseFloat(milsString, 32)
				if err != nil {
					fmt.Println("Please enter a valid number. Run matrixTool setMils --help for more information.")
					return
				}
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
	setMilsCmd.PersistentFlags().StringVarP(&milsString, "milsString", "", "", "[Required] Set mils of matrix pricing desired")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setMilsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setMilsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
