package excel

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

const mainSheetName = "Daily Matrix Price For All Term"

func ReadExcelFile(filePath string) {
	workbook, err := excelize.OpenFile(filePath)
	cobra.CheckErr(err)

	defer func(workbook *excelize.File) {
		err := workbook.Close()
		cobra.CheckErr(err)
	}(workbook)

	var style excelize.Style
	theStyle := &style
	(theStyle).NumFmt = 17
	var styleId int
	styleId, err = workbook.NewStyle(theStyle)
	err = workbook.SetColStyle(mainSheetName, "A", styleId)
	cobra.CheckErr(err)
	rows, err := workbook.GetRows(mainSheetName)
	cobra.CheckErr(err)
	//fmt.Println("asap")
	//dbModify.ProcessRow(rows[53])
	for _, row := range rows[53:134] {
		dbModify.ProcessRow(row)
	}
	fmt.Println()

	//db, openErr := sql.Open("sqlite", "./data.db")
	//cobra.CheckErr(openErr)
	//
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	cobra.CheckErr(err)
	//}(db)
	//
	////var param interface{}
	////param = true
	//query := `SELECT * FROM matrix WHERE contract_start = ? `
	//query += `AND billing_method != ?`
	//row, err := db.Query(query, "Jul-23", "Dual")
	//cobra.CheckErr(err)
	//entry := dbModify.MatrixEntry{}
	//defer func(row *sql.Rows) {
	//	err := row.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}(row)
	//for row.Next() {
	//	err := row.Scan(&entry.Id, &entry.ContractStart, &entry.State, &entry.Util, &entry.Zone, &entry.RateCode, &entry.ProductOption, &entry.BillingMethod, &entry.Term, &entry.UsageLower, &entry.UsageMiddle, &entry.UsageUpper)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(entry)
	//}

	entries := dbModify.GetFilteredEntries()
	fmt.Println()
	for _, entry := range entries {
		fmt.Printf("%+v\n", entry)
	}
}

//func WriteReport(filePath string)
