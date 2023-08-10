/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"regexp"

	"github.com/spf13/cobra"
)

// setTermsCmd represents the setTerms command
var (
	terms string

	setTermsCmd = &cobra.Command{
		Use:   "setTerms",
		Short: "Set the parameter for terms",
		Long: `Terms entered will be updated to inputted parameter
Any Reports generated will be under this terms
Please enter the terms as "[]" or in a 4ct array format: "[12,24,36,48]", "[]"
Note: MUST BE a 4 count array - enter duplicate values if needed: "[12,12,14,48]"
Valid input: matrixTool setTerms "[12,12,36,36]"
Invalid input: matrixTool setTerms "[12,36]"`,
		Run: func(cmd *cobra.Command, args []string) {
			// check if args exist
			if terms == "" && len(args) != 0 {
				terms = args[0]
				// check if input matches format
				regex := regexp.MustCompile("^[[0-9]+,[0-9]+,[0-9]+,[0-9]+\\]$")
				res := regex.MatchString(terms)
				if terms != "[]" {
					if !res {
						fmt.Println("Please specify terms with the correct format. Run 'matrixTool setTerms --help' for more information.")
						return
					}
				}
				fmt.Println("The entered terms code is: " + terms + "\n")
			} else {
				fmt.Println("Please specify terms. Run 'matrixTool setTerms --help' for more information.")
				return
			}
			// convert string to int slice & update parameters
			var termsIntSlice []int
			err := ffjson.Unmarshal([]byte(terms), &termsIntSlice)
			cobra.CheckErr(err)
			dbModify.SetTerms(termsIntSlice)
		},
	}
)

func init() {
	rootCmd.AddCommand(setTermsCmd)
	rootCmd.PersistentFlags().StringVarP(&terms, "terms", "", "", "[Required] Set terms of matrix pricing desired")
}
