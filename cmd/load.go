/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/rh5661/matrixTool/pkg/excel"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var (
	filePath string

	loadCmd = &cobra.Command{
		Use:   "load",
		Short: "Loads matrix excel file into database",
		Long: `Excel filepath inserted will be parsed and injected into database
Please specify the absolute filepath from the root (Example: C:\)
NOTE: Filepath can be acquired by right clicking the file, Ctrl+Shift+C, or dragging the file into the terminal window to paste the path
IMPORTANT: Make sure the Daily Matrix Price Report sheet is named: "Daily Matrix Price For All Term"
Example usage:
matrixTool load "C:\Users\Robert\Downloads\Daily Matrix Price For All Term.xlsx"`,
		Run: func(cmd *cobra.Command, args []string) {
			// check if args provided
			if filePath == "" && len(args) != 0 {
				filePath = args[0]
				fmt.Println("The entered filePath is: " + filePath + "\n")
			} else {
				fmt.Println("Please specify a filepath. Run 'matrixTool load --help' for more information.")
				return
			}

			// update parameters and load data from Excel file into db
			dbModify.SetFilePath(filePath)
			excel.ReadExcelFile(filePath)
			fmt.Println("\nFinished loading")
		},
	}
)

func init() {
	rootCmd.AddCommand(loadCmd)
	loadCmd.PersistentFlags().StringVarP(&filePath, "filePath", "", "", "[Required] Set filepath of matrix pricing Excel file to be parsed")
}
