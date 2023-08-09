/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

Example usage:
matrixTool generate`,
	Run: func(cmd *cobra.Command, args []string) {
		parameters := dbModify.ReadJson()
		filePathExcel := parameters.FilePath
		if filePathExcel != "" {
			fmt.Println("The filePath of the excel file is: " + filePathExcel + "\n")
		} else {
			fmt.Println("Please specify a filepath with load. Run 'matrixTool load --help' for more information.")
			return
		}

		entries := dbModify.GetFilteredEntries()
		dbModify.InsertMargin(entries, parameters.Mils)
		fmt.Println()
		for _, entry := range entries {
			fmt.Printf("%+v\n", entry)
		}
		excel.WriteReport(filePathExcel, entries)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	//loadCmd.PersistentFlags().StringVarP(&filePathExcel, "filePathExcel", "", "", "[Required] Set filepath of matrix pricing Excel file to be parsed")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
