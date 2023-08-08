package excel

import (
	"fmt"
	"github.com/rh5661/matrixTool/pkg/dbModify"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"
	"time"
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
	dbModify.ReInitializeDatabase()
	for _, row := range rows[53:134] { // CHANGE THIS TO ALL ROWS
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

}

func WriteReport(filePath string, entries []dbModify.MatrixEntry) {
	workbook, err := excelize.OpenFile(filePath)
	cobra.CheckErr(err)

	defer func(workbook *excelize.File) {
		err := workbook.Close()
		cobra.CheckErr(err)
	}(workbook)

	userParameters := dbModify.ReadJson()
	datetime := time.Now()
	sheetName := fmt.Sprintf("%s%s%s-%s-%s", userParameters.StartDate, userParameters.Util, strconv.Itoa(datetime.Hour()), strconv.Itoa(datetime.Minute()), strconv.Itoa(datetime.Second()))
	_, err = workbook.NewSheet(sheetName)
	if err != nil {
		fmt.Println(sheetName)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Sheet created: " + sheetName)

	err = workbook.SetColWidth(sheetName, "A", "A", 9)
	err = workbook.SetColWidth(sheetName, "B", "B", 9)
	err = workbook.SetColWidth(sheetName, "C", "C", 9)
	err = workbook.SetColWidth(sheetName, "D", "D", 9)
	err = workbook.SetColWidth(sheetName, "E", "E", 23)
	err = workbook.SetColWidth(sheetName, "F", "F", 9)
	err = workbook.SetColWidth(sheetName, "G", "G", 9)
	err = workbook.SetColWidth(sheetName, "H", "H", 9)
	err = workbook.SetColWidth(sheetName, "I", "I", 9)
	err = workbook.SetColWidth(sheetName, "J", "J", 9)
	err = workbook.SetColWidth(sheetName, "K", "K", 9)
	err = workbook.SetRowHeight(sheetName, 1, 45)
	fmt.Println("Sizing set...")

	borders := []excelize.Border{{Type: "top", Color: "#000000", Style: 1}, {Type: "left", Color: "#000000", Style: 1}, {Type: "right", Color: "#000000", Style: 1}, {Type: "bottom", Color: "#000000", Style: 1}}
	style, err := workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 11, Family: "Calibri"}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}, Border: borders})
	err = workbook.SetColStyle(sheetName, "A:K", style)
	style, err = workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 11, Bold: true, Color: "#FFFFFF", Family: "Calibri"}, Fill: excelize.Fill{Type: "pattern", Color: []string{"#00008B"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}, Border: borders})
	err = workbook.SetCellStyle(sheetName, "A1", "K1", style)
	fmt.Println("Styles set...")

	headers := []string{"Contract Start Month", "State", "Utility", "Zone", "Rate Code(s)", "Product Special Options", "Billing Method", "Term", "0 - 49", "50 - 299", "300 - 1099"}
	err = workbook.SetSheetRow(sheetName, "A1", &headers)
	fmt.Println("Headers inserted...")

	startRowIndex := 2
	for _, entry := range entries {
		err = workbook.SetRowHeight(sheetName, startRowIndex, 90)
		startCell := "A" + strconv.Itoa(startRowIndex)
		entrySlice := []string{entry.ContractStart, entry.State, entry.Util, entry.Zone, entry.RateCode, entry.ProductOption, entry.BillingMethod, strconv.Itoa(entry.Term), fmt.Sprintf("%.5f", entry.UsageLower), fmt.Sprintf("%.5f", entry.UsageMiddle), fmt.Sprintf("%.5f", entry.UsageUpper)}
		err = workbook.SetSheetRow(sheetName, startCell, &entrySlice)
		fmt.Println("Entry inserted: " + startCell)
		startRowIndex++
	}

	date, err := workbook.GetCellValue(mainSheetName, "A3")
	date = strings.ReplaceAll(date, "as of ", "")
	params := dbModify.ReadJson()
	var startDate string
	switch params.StartDate[:3] {
	case "Jan":
		startDate = "January"
	case "Feb":
		startDate = "February"
	case "Mar":
		startDate = "March"
	case "Apr":
		startDate = "April"
	case "May":
		startDate = "May"
	case "Jun":
		startDate = "June"
	case "Jul":
		startDate = "July"
	case "Aug":
		startDate = "August"
	case "Sep":
		startDate = "September"
	case "Oct":
		startDate = "October"
	case "Nov":
		startDate = "November"
	case "Dec":
		startDate = "December"
	default:
		startDate = "ERROR"
		fmt.Println("ERROR: READ START DATE")
	}
	startDate += " 20" + params.StartDate[4:]
	infoText := fmt.Sprintf("%s %s Start (%s)", params.Util, startDate, date)
	infoStartCell := "A" + strconv.Itoa(startRowIndex)
	infoEndCell := "K" + strconv.Itoa(startRowIndex+3)
	err = workbook.MergeCell(sheetName, infoStartCell, infoEndCell)
	style, err = workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 28, Bold: true}})
	err = workbook.SetCellStyle(sheetName, infoStartCell, infoEndCell, style)
	err = workbook.SetSheetRow(sheetName, infoStartCell, &[]interface{}{infoText})
	fmt.Println("Info text inserted...")

	err = workbook.Save()
	cobra.CheckErr(err)
	fmt.Println("File Saved at: " + filePath)

}
