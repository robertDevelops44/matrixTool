/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

// showUtilsCmd represents the showUtils command
var showUtilsCmd = &cobra.Command{
	Use:   "showUtils",
	Short: "Displays all valid utilities with their code & full name",
	Long: `Displays the utility codes valid for input as a parameter
The full name of each utility is provided alongside each shorthand code
To change the parameter for utility code, please run the following for more information:
matrixTool setUtil --help
`,
	Run: func(cmd *cobra.Command, args []string) {
		// make array to store all utility code/name pairs
		utilCodes := make([]string, len(dbModify.UtilCodes))
		fmt.Println("Valid utility codes: ")
		for code, utility := range dbModify.UtilCodes {
			pair := fmt.Sprintf("%s : %s\n", code, utility)
			utilCodes = append(utilCodes, pair)
		}
		// sort and print
		sort.Strings(utilCodes)
		printString := fmt.Sprintf("%v", utilCodes)
		printString = strings.TrimSpace(printString[1 : len(printString)-1])
		fmt.Println(" " + printString + "\n")
	},
}

func init() {
	rootCmd.AddCommand(showUtilsCmd)
}
