/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:     "matrixTool",
	Version: "1.0.0",
	Short:   "Tool for producing filtered matrix pricing",
	Long:    `matrixTool is a CLI application that takes an Excel file with matrix pricing and produces filtered pricing with custom broker fee and adjustments.`,
}

func Execute() {
	// init db
	dbModify.InitializeDatabase()

	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
