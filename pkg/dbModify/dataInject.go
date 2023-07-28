package dbModify

import "fmt"

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

	fmt.Printf(`Contract Start Month: %s
State: %s
Utility: %s
Zone: %s
Rate Code(s): %s
Product Special Options: %s
Billing Method: %s
Term: %s
0-49: %s
50-299: %s
300-1099: %s`, contractStart, state, utility, zone, rateCodes, productOptions, billingMethod, term, usageLower, usageMiddle, usageUpper)
	return true
}
