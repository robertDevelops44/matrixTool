/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/excel"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var (
	filePath string

	loadCmd = &cobra.Command{
		Use:   "load",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if filePath == "" && len(args) != 0 {
				filePath = args[0]
				fmt.Println("The entered filePath is: " + filePath + "\n")
			} else {
				fmt.Println("Please specify a filepath. Run 'matrixTool load --help' for more information.")
				return
			}

			excel.ReadExcelFile(filePath)
			fmt.Println("\nFinished loading")
		},
	}
)

func init() {
	rootCmd.AddCommand(loadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	loadCmd.PersistentFlags().StringVarP(&filePath, "filePath", "", "", "[Required] Set filepath of matrix pricing Excel file to be parsed")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
