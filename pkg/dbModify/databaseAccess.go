package dbModify

import (
	"database/sql"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/spf13/cobra"
	"math"
	"os"
	"strconv"
)

// MatrixEntry
//
//	@Description: struct to store matrix pricing entry
type MatrixEntry struct {
	Id                                                                       int
	ContractStart, State, Util, Zone, RateCode, ProductOption, BillingMethod string
	Term                                                                     int
	UsageLower, UsageMiddle, UsageUpper                                      float32
}

// QueryParameters
//
//	@Description: struct to store parameters
type QueryParameters struct {
	FilePath    string  `json:"filePath"`    // file path of Excel file
	Mils        float32 `json:"mils"`        // broker fee for report
	StartDate   string  `json:"startDate"`   // contract start month/year
	Util        string  `json:"util"`        // utility code
	DualBilling bool    `json:"dualBilling"` // include dual billing
	Terms       []int   `json:"terms"`       // terms to include
}

// UserParameters
//
//	@Description: var to store parameters with each instance
var UserParameters = QueryParameters{}

// UtilCodes
//
//	@Description: map of all valid utility codes and names
var UtilCodes = map[string]string{
	"AECO":    "Atlantic City Electric Company",
	"APS":     "Potomac Edison - Allegheny Power",
	"BED":     "Eversource - Boston Edison",
	"BGE":     "Baltimore Gas and Electric",
	"CAMB":    "Eversource - Cambridge Electric",
	"CEI":     "Cleveland Electric",
	"CGE":     "Duke Energy",
	"CHGE":    "Central Hudson",
	"CILCO":   "Ameren Rate Zone II - CILCO",
	"CIPS":    "Ameren Rate Zone I - CIPS",
	"CMP":     "Central Maine Power",
	"COME":    "Eversource - Commonwealth Electric",
	"COMED":   "Commonwealth Edison",
	"CONE":    "Consolidated Edison",
	"CS":      "AEP - CS",
	"DELM":    "Conectiv Delmarva",
	"DELMDE":  "Delmarva",
	"DLCO":    "Duquesne Light Company",
	"DPL":     "Dayton Power and Light",
	"FGE":     "Unitil - Fitchburg Gas and Electric",
	"GSECO":   "Granite State Electric Co (Liberty Utilities)",
	"ILPWR":   "Ameren Rate Zone III - IP",
	"JCPL":    "Jersey Central Power & Light Company",
	"METED":   "Metropolitan Edison Company",
	"MSEL":    "National Grid - Massachusetts Electric Company",
	"MWST":    "Eversource - Western Massachusetts Electric",
	"NECO":    "National Grid - Narragansett Electric",
	"NHEC":    "New Hampshire Electric Co",
	"NIMO":    "National Grid - Niagara Mohawk",
	"NYOR":    "Orange and Rockland",
	"NYSEG":   "New York State Electric & Gas",
	"OE":      "Ohio Edison",
	"OPCO":    "AEP - OP",
	"PECO":    "PECO Energy",
	"PENELEC": "Pennsylvania Electric Company",
	"PEPCO":   "Potomac Electric Power Company",
	"PP":      "Pennsylvania Power Company",
	"PPL":     "Pennsylvania Power and Light, Inc.",
	"PSEG":    "Public Service Electric and Gas Company",
	"PSNH":    "Public Service Of New Hampshire",
	"RECO":    "Rockland Electric Company",
	"RGE":     "Rochester Gas & Electric",
	"TE":      "Toledo Edison",
	"UNITIL":  "Unitil Energy Systems",
	"WPP":     "Allegheny Power WPP",
}

// parametersFilePath dataSourcePath
//
//	@Description: file paths for persistent data storage
const parametersFilePath = "./parameters.json"
const dataSourcePath = "./data.db"

// dbInitSQL
//
//	@Description: sql statement for initializing the database
const dbInitSQL string = `

CREATE TABLE IF NOT EXISTS matrix (
	id 					INTEGER PRIMARY KEY,
	contract_start 		DATE NOT NULL,
	state_code 			TEXT NOT NULL,
	util_code			TEXT NOT NULL,
	util_zone			TEXT NOT NULL,
	util_rate_code		TEXT,
	product_option		TEXT NOT NULL,
	billing_method 		TEXT NOT NULL,
	contract_term		INTEGER NOT NULL,
	usage_lower			FLOAT NOT NULL,
	usage_middle 		FLOAT NOT NULL,
	usage_upper 		FLOAT NOT NULL
);

`

// ProcessRow
//
//	@Description: inserts Excel row of pricing into db
//	@param row
//	@return bool
func ProcessRow(row []string) bool {
	// format row
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

	// open db
	db, openErr := sql.Open("sqlite", dataSourcePath)
	cobra.CheckErr(openErr)

	defer func(db *sql.DB) {
		err := db.Close()
		cobra.CheckErr(err)
	}(db)

	//insert into db
	insertSQL := `INSERT INTO matrix (contract_start, state_code, util_code, util_zone, util_rate_code, product_option, billing_method, contract_term, usage_lower, usage_middle, usage_upper) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?);`
	stmt, err := db.Prepare(insertSQL)
	cobra.CheckErr(err)
	res, err := stmt.Exec(contractStart, state, utility, zone, rateCodes, productOptions, billingMethod, term, usageLower, usageMiddle, usageUpper)
	cobra.CheckErr(err)

	id, err := res.LastInsertId()
	cobra.CheckErr(err)

	fmt.Println("Id inserted: " + strconv.FormatInt(id, 10))
	return true
}

// GetFilteredEntries
//
//	@Description: retrieve for all pricing entries in db based on parameters and convert to MatrixEntry structs
//	@return []MatrixEntry
func GetFilteredEntries() []MatrixEntry {
	fmt.Println(`Getting filtered entries...`)

	// load params
	userParameters := ReadJson()
	LoadParameters(userParameters)

	// open db
	db, openErr := sql.Open("sqlite", dataSourcePath)
	cobra.CheckErr(openErr)

	defer func(db *sql.DB) {
		err := db.Close()
		cobra.CheckErr(err)
	}(db)

	// construct query string
	numParams := 0
	var paramsToInsert []string
	querySQL := `SELECT * FROM matrix WHERE `

	var rows *sql.Rows
	var err error

	fmt.Println("\nStart Date: " + UserParameters.StartDate)
	if UserParameters.StartDate != "" {
		querySQL += `contract_start = ? AND `
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

	// go through rows and turn into Matrix Entry struct
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

// InitializeDatabase
//
//	@Description: initialize db - create table
func InitializeDatabase() {
	db, openErr := sql.Open("sqlite", dataSourcePath)
	cobra.CheckErr(openErr)

	defer func(db *sql.DB) {
		err := db.Close()
		cobra.CheckErr(err)
	}(db)
	_, err := db.Exec(dbInitSQL)
	cobra.CheckErr(err)
}

// ReInitializeDatabase
//
//	@Description: initialize db with drop table statement
func ReInitializeDatabase() {
	db, openErr := sql.Open("sqlite", dataSourcePath)
	cobra.CheckErr(openErr)

	defer func(db *sql.DB) {
		err := db.Close()
		cobra.CheckErr(err)
	}(db)
	dropInitSQL := `DROP TABLE IF EXISTS matrix;` + dbInitSQL
	_, err := db.Exec(dropInitSQL)
	cobra.CheckErr(err)
}

// LoadParameters
//
//	@Description: load parameters into variable UserParameters
//	@param newParameters
func LoadParameters(newParameters QueryParameters) {
	UserParameters = newParameters
	PrintParameters()
}

// PrintParameters
//
//	@Description: print formatted parameters
func PrintParameters() {
	fmt.Printf("User Parameters:\n%v\n", UserParameters)
}

// ReadJson
//
//	@Description: read parameters json file and convert to QueryParameters struct
//	@return QueryParameters
func ReadJson() QueryParameters {
	data, err := os.ReadFile(parametersFilePath)
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

// writeJson
//
//	@Description: write/update UserParameters into parameters json file
func writeJson() {
	data, err := easyjson.Marshal(UserParameters)
	cobra.CheckErr(err)

	err = os.WriteFile("parameters.json", data, os.ModePerm)
	cobra.CheckErr(err)

}

// SetFilePath
//
//	@Description: updates FilePath in UserParameters - filepath of Excel file
//	@param filePath
func SetFilePath(filePath string) {
	parameters := ReadJson()
	fmt.Println("Old:")
	LoadParameters(parameters)
	(UserParameters).FilePath = filePath
	writeJson()
	fmt.Println("New:")
	PrintParameters()
}

// SetMils
//
//	@Description: updates Mils in UserParameters - broker fee
//	@param mils
func SetMils(mils float32) {
	parameters := ReadJson()
	fmt.Println("Old:")
	LoadParameters(parameters)
	(UserParameters).Mils = mils
	writeJson()
	fmt.Println("New:")
	PrintParameters()
}

// SetStartDate
//
//	@Description: updates StartDate in UserParameters - contract start month/year
//	@param startDate
func SetStartDate(startDate string) {
	parameters := ReadJson()
	fmt.Println("Old:")
	LoadParameters(parameters)
	(UserParameters).StartDate = startDate
	writeJson()
	fmt.Println("New:")
	PrintParameters()
}

// SetUtil
//
//	@Description: update Util in UserParameters - utility code
//	@param util
func SetUtil(util string) {
	parameters := ReadJson()
	fmt.Println("Old:")
	LoadParameters(parameters)
	(UserParameters).Util = util
	writeJson()
	fmt.Println("New:")
	PrintParameters()
}

// SetDualBilling
//
//	@Description: update DualBilling in UserParameters - include dual billing
//	@param dualBilling
func SetDualBilling(dualBilling bool) {
	parameters := ReadJson()
	fmt.Println("Old:")
	LoadParameters(parameters)
	(UserParameters).DualBilling = dualBilling
	writeJson()
	fmt.Println("New:")
	PrintParameters()
}

// SetTerms
//
//	@Description: update Terms in UserParameters - terms to include
//	@param terms
func SetTerms(terms []int) {
	parameters := ReadJson()
	fmt.Println("Old:")
	LoadParameters(parameters)
	(UserParameters).Terms = terms
	writeJson()
	fmt.Println("New:")
	PrintParameters()
}

// InsertMargin
//
//	@Description: adds margin to each entry in a []MatrixEntry and adjusts price accordingly
//	@param entries
//	@param mils
func InsertMargin(entries []MatrixEntry, mils float32) {
	// update all 3 prices in each entry
	for i, entry := range entries {
		entryPtr := &entries[i]
		entryPtr.UsageLower = calculatePricing(entry.UsageLower, mils)
		entryPtr.UsageMiddle = calculatePricing(entry.UsageMiddle, mils)
		entryPtr.UsageUpper = calculatePricing(entry.UsageUpper, mils)
	}
}

// calculatePricing
//
//	@Description: calculates and produces price given broker fee, adjusts if needed
//	@param initialPrice
//	@param mils
//	@return float32
func calculatePricing(initialPrice float32, mils float32) float32 {
	price := (initialPrice + mils) / 1000
	if initialPrice == 0 {
		price = 0
	} else if math.Mod(float64(price), 0.01) <= 0.002 { // if the 1000th place is a 0 or 1 - 0.12053, 0.09022
		// round down to 0.xx998 	- 	0.12053 -> 0.11998
		price = float32(math.Floor(float64(price*100)) / 100)
		price -= 0.01
		price += .00998
	}
	return price
}

// GetUtilByCode
//
//	@Description: get utility name by shorthand code
//	@param code
//	@return string
func GetUtilByCode(code string) string {
	return UtilCodes[code]
}
