/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utilCodes := make([]string, len(dbModify.UtilCodes))
		fmt.Println("Valid utility codes: ")
		for code, utility := range dbModify.UtilCodes {
			pair := fmt.Sprintf("%s : %s\n", code, utility)
			utilCodes = append(utilCodes, pair)
		}
		sort.Strings(utilCodes)
		printString := fmt.Sprintf("%v", utilCodes)
		printString = strings.TrimSpace(printString[1 : len(printString)-1])
		fmt.Println(printString + "\n")
	},
}

func init() {
	rootCmd.AddCommand(showUtilsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showUtilsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showUtilsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
