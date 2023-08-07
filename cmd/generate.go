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
var (
	filePathExcel string

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if filePathExcel == "" && len(args) != 0 {
				filePathExcel = args[0]
				fmt.Println("The entered filePath is: " + filePathExcel + "\n")
			} else {
				fmt.Println("Please specify a filepath. Run 'matrixTool generate --help' for more information.")
				return
			}

			excel.ReadExcelFile(filePathExcel)
			fmt.Println("\nFinished loading")
			entries := dbModify.GetFilteredEntries()
			fmt.Println()
			for _, entry := range entries {
				fmt.Printf("%+v\n", entry)
			}
			excel.WriteReport(filePathExcel, entries)
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	loadCmd.PersistentFlags().StringVarP(&filePathExcel, "filePathExcel", "", "", "[Required] Set filepath of matrix pricing Excel file to be parsed")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
