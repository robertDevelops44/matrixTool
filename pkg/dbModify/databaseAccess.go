package dbModify

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

type QueryParameters struct {
	StartDate   string `json:"startDate"`
	Util        string `json:"util"`
	DualBilling bool   `json:"dualBilling"`
	Terms       []int  `json:"terms"`
}

var userParameters = QueryParameters{}

func ProcessRow(row []string) bool {

	var contractStart, state, utility, zone, rateCodes, productOptions, billingMethod, term, usageLower, usageMiddle, usageUpper string

	contractStart = row[0]
	state = row[1]
	utility = row[2]
	zone = row[3]
	rateCodes = row[4]
	productOptions = row[5]
	billingMethod = row[6]
	term = row[7]
	usageLower = row[8]
	usageMiddle = row[9]
	usageUpper = row[10]

	db, openErr := sql.Open("sqlite", "./data.db")
	cobra.CheckErr(openErr)

	defer func(db *sql.DB) {
		err := db.Close()
		cobra.CheckErr(err)
	}(db)
	stmt, err := db.Prepare(`INSERT INTO matrix (contract_start, state_code_id, util_code_id, util_zone, util_rate_code, product_option, billing_method, contract_term, usage_lower, usage_middle, usage_upper) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?);`)
	cobra.CheckErr(err)
	res, err := stmt.Exec(contractStart, state, utility, zone, rateCodes, productOptions, billingMethod, term, usageLower, usageMiddle, usageUpper)
	cobra.CheckErr(err)

	id, err := res.LastInsertId()
	cobra.CheckErr(err)

	fmt.Println("Id inserted: " + strconv.FormatInt(id, 10))

	//	fmt.Printf(`Contract Start Month: %s
	//State: %s
	//Utility: %s
	//Zone: %s
	//Rate Code(s): %s
	//Product Special Options: %s
	//Billing Method: %s
	//Term: %s
	//0-49: %s
	//50-299: %s
	//300-1099: %s`, contractStart, state, utility, zone, rateCodes, productOptions, billingMethod, term, usageLower, usageMiddle, usageUpper)
	return true
}

func LoadParameters(newParameters QueryParameters) {
	//var memAdd = &userParameters
	//*memAdd = newParameters
	userParameters := newParameters
	fmt.Println("user parameters records: ")
	fmt.Println(userParameters)
}

func SetStartDate(startDate string) {
	(userParameters).StartDate = startDate
}

func SetUtil(util string) {
	(userParameters).Util = util
}

func SetDualBilling(dualBilling bool) {
	(userParameters).DualBilling = dualBilling
}

func SetTerms(terms []int) {
	(userParameters).Terms = terms
}
