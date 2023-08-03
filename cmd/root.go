/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "matrixTool",
	Version: "1.0.0",
	Short:   "Tool for producing filtered matrix pricing",
	Long:    `REPLACE WITH COMPLETE DESC.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// open db
	dbModify.InitializeDatabase()

	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.MatrixTool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
