package dbModify

import (
	"database/sql"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

type MatrixEntry struct {
	Id                                                                       int
	ContractStart, State, Util, Zone, RateCode, ProductOption, BillingMethod string
	Term                                                                     int
	UsageLower, UsageMiddle, UsageUpper                                      float32
}

type QueryParameters struct {
	StartDate   string `json:"startDate"`
	Util        string `json:"util"`
	DualBilling bool   `json:"dualBilling"`
	Terms       []int  `json:"terms"`
}

var UserParameters = QueryParameters{}

const parametersFilePath = "./parameters.json"

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

	insertSQL := `INSERT INTO matrix (contract_start, state_code, util_code, util_zone, util_rate_code, product_option, billing_method, contract_term, usage_lower, usage_middle, usage_upper) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?);`
	stmt, err := db.Prepare(insertSQL)
	cobra.CheckErr(err)
	res, err := stmt.Exec(contractStart, state, utility, zone, rateCodes, productOptions, billingMethod, term, usageLower, usageMiddle, usageUpper)
	cobra.CheckErr(err)

	id, err := res.LastInsertId()
	cobra.CheckErr(err)

	fmt.Println("Id inserted: " + strconv.FormatInt(id, 10))

	//fmt.Printf(`Contract Start Month: %s
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

func GetFilteredEntries() []MatrixEntry {
	fmt.Println(`Getting filtered entries...`)

	userParameters := ReadJson(parametersFilePath)
	LoadParameters(userParameters)

	db, openErr := sql.Open("sqlite", "./data.db")
	cobra.CheckErr(openErr)

	defer func(db *sql.DB) {
		err := db.Close()
		cobra.CheckErr(err)
	}(db)

	numParams := 0
	var paramsToInsert []string
	querySQL := `SELECT * FROM matrix WHERE `

	var rows *sql.Rows
	var err error

	fmt.Println("Start Date: " + UserParameters.StartDate)
	if UserParameters.StartDate != "" {
		querySQL += `contract_start = ? AND `
		//startParam = UserParameters.StartDate
		paramsToInsert = append(paramsToInsert, UserParameters.StartDate)
		numParams++
	}
	fmt.Println("Util: " + UserParameters.Util)
	if UserParameters.Util != "" {
		querySQL += `util_code = ? AND `
		paramsToInsert = append(paramsToInsert, UserParameters.Util)
		numParams++
	}
	fmt.Println("Dual Billing: " + strconv.FormatBool(UserParameters.DualBilling))
	if !(UserParameters.DualBilling) {
		querySQL += `billing_method != 'Dual' AND `
	}
	termsString, err := ffjson.Marshal(UserParameters.Terms)
	fmt.Println("Terms: " + string(termsString))
	if len(UserParameters.Terms) == 4 {
		querySQL += `contract_term IN (?,?,?,?);`
		switch numParams {
		case 0:
			rows, err = db.Query(querySQL, (UserParameters.Terms)[0], (UserParameters.Terms)[1], (UserParameters.Terms)[2], (UserParameters.Terms)[3])
			cobra.CheckErr(err)
		case 1:
			rows, err = db.Query(querySQL, paramsToInsert[0], (UserParameters.Terms)[0], (UserParameters.Terms)[1], (UserParameters.Terms)[2], (UserParameters.Terms)[3])
			cobra.CheckErr(err)
		case 2:
			rows, err = db.Query(querySQL, paramsToInsert[0], paramsToInsert[1], (UserParameters.Terms)[0], (UserParameters.Terms)[1], (UserParameters.Terms)[2], (UserParameters.Terms)[3])
			cobra.CheckErr(err)
		default:
			fmt.Println("error")
		}
	} else {
		querySQL += `contract_term IN (contract_term);`
		switch numParams {
		case 0:
			rows, err = db.Query(querySQL)
			cobra.CheckErr(err)
		case 1:
			rows, err = db.Query(querySQL, paramsToInsert[0])
			cobra.CheckErr(err)
		case 2:
			rows, err = db.Query(querySQL, paramsToInsert[1])
			cobra.CheckErr(err)
		default:
			fmt.Println("error")
		}
	}

	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	var entries []MatrixEntry

	for rows.Next() {
		entry := MatrixEntry{}
		err := rows.Scan(&entry.Id, &entry.ContractStart, &entry.State, &entry.Util, &entry.Zone, &entry.RateCode, &entry.ProductOption, &entry.BillingMethod, &entry.Term, &entry.UsageLower, &entry.UsageMiddle, &entry.UsageUpper)
		if err != nil {
			fmt.Println(err)
		}
		entries = append(entries, entry)
	}

	return entries

}

func LoadParameters(newParameters QueryParameters) {
	//var memAdd = &UserParameters
	//*memAdd = newParameters
	UserParameters = newParameters
	fmt.Println("User Parameters processed: ")
	fmt.Println(UserParameters)
}

func ReadJson(filePath string) QueryParameters {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return QueryParameters{}
	}
	defaultParameters := &(QueryParameters{})
	err = easyjson.Unmarshal(data, defaultParameters)
	if err != nil {
		fmt.Println(err)
		return QueryParameters{}
	}
	return *defaultParameters
}

func SetStartDate(startDate string) {
	(UserParameters).StartDate = startDate
}

func SetUtil(util string) {
	(UserParameters).Util = util
}

func SetDualBilling(dualBilling bool) {
	(UserParameters).DualBilling = dualBilling
}

func SetTerms(terms []int) {
	(UserParameters).Terms = terms
}
