/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
)

// showParametersCmd represents the showParameters command
var showParametersCmd = &cobra.Command{
	Use:   "showParameters",
	Short: "Displays current parameters entered",
	Long: `Displays the parameters that will be used if matrixTool generate is called
To change these parameters, please run the following for more information:
matrixTool load --help
matrixTool setStart --help
matrixTool setUtil --help
matrixTool setDualBilling --help
matrixTool setTerms --help
matrixTool setMils --help`,
	Run: func(cmd *cobra.Command, args []string) {
		// get parameters, format, print
		params := dbModify.ReadJson()
		paramsMarshalled, err := json.MarshalIndent(params, "", "	")
		cobra.CheckErr(err)
		fmt.Printf("User Parameters:\n%s", paramsMarshalled)
	},
}

func init() {
	rootCmd.AddCommand(showParametersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showParametersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showParametersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
