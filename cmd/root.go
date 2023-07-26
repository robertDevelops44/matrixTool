/*
Copyright © 2023 ROBERT HUANG
*/
package cmd

import (
	"database/sql"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

const dbInit string = `
DROP TABLE IF EXISTS matrix;

CREATE TABLE IF NOT EXISTS matrix (
	id 					INTEGER NOT NULL PRIMARY KEY,
	contract_start 		DATE NOT NULL,
	state_code_id 		INTEGER NOT NULL,
	util_code_id		INTEGER NOT NULL,
	util_zone			TEXT NOT NULL,
	util_rate_code		TEXT,
	product_option		TEXT NOT NULL,
	billing_method 		TEXT NOT NULL,
	contract_term		INTEGER NOT NULL,
	usage_lower			DOUBLE NOT NULL,
	usage_middle 		DOUBLE NOT NULL,
	usage_upper 		DOUBLE NOT NULL
);

CREATE TABLE IF NOT EXISTS state_codes (
    id				INTEGER NOT NULL PRIMARY KEY,
    state_code		TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS util_codes (
    id				INTEGER NOT NULL PRIMARY KEY,
    util_code		TEXT NOT NULL,
	util_name		TEXT NOT NULL
);

INSERT INTO state_codes (state_code) VALUES 
	("DE"),
	("IL"),
	("MA"),
	("MD"),
	("ME"),
	("NH"),
	("NJ"),
	("OH"),
	("PA"),
	("RI");

INSERT INTO util_codes (util_code,util_name) VALUES
	("AECO",	"Atlantic City Electric Company"),			
	("APS",		"Potomac Edison - Allegheny Power"),			
	("BED",		"Eversource - Boston Edison"), 			
	("BGE",		"Baltimore Gas and Electric"),			
	("CAMB",	"Eversource - Cambridge Electric"), 			
	("CEI",		"Cleveland Electric"),			
	("CGE",		"Duke Energy"),			
	("CHGE",	"Central Hudson"),			
	("CILCO",	"Ameren Rate Zone II - CILCO"),			
	("CIPS",	"Ameren Rate Zone I - CIPS"),			
	("CMP",		"Central Maine Power"),			
	("COME",	"Eversource - Commonwealth Electric"),			
	("COMED",	"Commonwealth Edison"),			
	("CONE",	"Consolidated Edison"),			
	("CS",		"AEP - CS"),			
	("DELM",	"Conectiv Delmarva"),			
	("DELMDE",	"Delmarva"),			
	("DLCO",	"Duquesne Light Company"),
	("DPL",		"Dayton Power and Light"),
	("FGE",		"Unitil - Fitchburg Gas and Electric"),			
	("GSECO",	"Granite State Electric Co (Liberty Utilities)"),			
	("ILPWR",	"Ameren Rate Zone III - IP"),			
	("JCPL",	"Jersey Central Power & Light Company"),			
	("METED",	"Metropolitan Edison Company"),			
	("MSEL",	"National Grid - Massachusetts Electric Company"),			
	("MWST",	"Eversource - Western Massachusetts Electric"),			
	("NECO",	"National Grid - Narragansett Electric"),			
	("NHEC",	"New Hampshire Electric Co"),			
	("NIMO",	"National Grid - Niagara Mohawk"),			
	("NYOR",	"Orange and Rockland"),			
	("NYSEG",	"New York State Electric & Gas"),			
	("OE",		"Ohio Edison"),		
	("OPCO",	"AEP - OP"),			
	("PECO",	"PECO Energy"),			
	("PENELEC",	"Pennsylvania Electric Company"),			
	("PEPCO",	"Potomac Electric Power Company"),			
	("PP",		"Pennsylvania Power Company"),			
	("PPL",		"Pennsylvania Power and Light, Inc."),			
	("PSEG",	"Public Service Electric and Gas Company"),			
	("PSNH",	"Public Service Of New Hampshire"),			
	("RECO",	"Rockland Electric Company"),			
	("RGE",		"Rochester Gas & Electric"),			
	("TE",		"Toledo Edison"),			
	("UNITIL",	"Unitil Energy Systems"),			
	("WPP",		"Allegheny Power WPP");

`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "MatrixTool",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// open db
	db, openErr := sql.Open("sqlite", "./data.db")
	cobra.CheckErr(openErr)

	defer func(db *sql.DB) {
		err := db.Close()
		cobra.CheckErr(err)
	}(db)

	_, createErr := db.Exec(dbInit)
	cobra.CheckErr(createErr)

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
