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

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Creates a new sheet with filtered pricing and margin",
	Long: `Edits the excel file at the filepath in parameters.json in the filePath 
Creates a new sheet with the name "{StartDate}{Util}{CurrentTime}"
The pricing will be filtered based on the entered parameters.
Margin will be inserted based on entered parameters and adjusted accordingly.
An info box will be generated below the pricing sheeting with Util,Start Date, Date of Matrix Used.
NOTE: Please make sure data is loaded into database before generating any reports
Run matrixTool load --help for more information
IMPORTANT: Excel file being modified must be closed while generating a report/pricing sheet
Make sure the Daily Matrix Price Report sheet is named: "Daily Matrix Price For All Term"

Example usage:
matrixTool generate`,
	Run: func(cmd *cobra.Command, args []string) {
		// get filepath of Excel file
		parameters := dbModify.ReadJson()
		filePathExcel := parameters.FilePath

		// make sure filepath is not blank
		if filePathExcel != "" {
			fmt.Println("The filePath of the excel file is: " + filePathExcel + "\n")
		} else {
			fmt.Println("Please specify a filepath with load. Run 'matrixTool load --help' for more information.")
			return
		}

		// fetch filtered pricing and add margin
		entries := dbModify.GetFilteredEntries()
		dbModify.InsertMargin(entries, parameters.Mils)
		fmt.Println()
		// print each pricing entry
		for _, entry := range entries {
			fmt.Printf("%+v\n", entry)
		}
		// create sheet in Excel file & populate with data
		excel.WriteReport(filePathExcel, entries)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
