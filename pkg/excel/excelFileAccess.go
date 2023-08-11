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

// mainSheetName default name for pricing sheet in Excel file
const mainSheetName = "Daily Matrix Price For All Term"

// ReadExcelFile
//
//	@Description: read Excel file, parse all entries, insert into db
//	@param filePath
func ReadExcelFile(filePath string) {
	// open Excel file
	workbook, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Printf("ERROR: %s! See matrixTool load --help for more information", err)
		os.Exit(1)
	}

	defer func(workbook *excelize.File) {
		err := workbook.Close()
		cobra.CheckErr(err)
	}(workbook)

	// re-style workbook for data retrieval
	var style excelize.Style
	theStyle := &style
	(theStyle).NumFmt = 17
	var styleId int
	styleId, err = workbook.NewStyle(theStyle)
	err = workbook.SetColStyle(mainSheetName, "A", styleId)
	cobra.CheckErr(err)
	rows, err := workbook.GetRows(mainSheetName)
	cobra.CheckErr(err)
	// reset db
	dbModify.ReInitializeDatabase()
	// insert processed Excel data
	for _, row := range rows[53:] { // CHANGE THIS TO ALL ROWS
		dbModify.ProcessRow(row)
	}
	fmt.Println()
}

// WriteReport
//
//	@Description: create new sheet in Excel file and generate pricing with broker fee
//	@param filePath
//	@param entries
func WriteReport(filePath string, entries []dbModify.MatrixEntry) {
	// open Excel file
	workbook, err := excelize.OpenFile(filePath)
	cobra.CheckErr(err)

	defer func(workbook *excelize.File) {
		err := workbook.Close()
		cobra.CheckErr(err)
	}(workbook)

	// get parameters, current time
	userParameters := dbModify.ReadJson()
	datetime := time.Now()
	// create new sheet with unique sheet name
	sheetName := fmt.Sprintf("%s%s%s-%s-%s", userParameters.StartDate, userParameters.Util, strconv.Itoa(datetime.Hour()), strconv.Itoa(datetime.Minute()), strconv.Itoa(datetime.Second()))
	_, err = workbook.NewSheet(sheetName)
	if err != nil {
		fmt.Println(sheetName)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Sheet created: " + sheetName)

	// set sheet sizing
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
	err = workbook.SetColWidth(sheetName, "L", "L", 50)
	err = workbook.SetColWidth(sheetName, "M", "M", 50)
	err = workbook.SetColWidth(sheetName, "N", "N", 50)
	err = workbook.SetColWidth(sheetName, "O", "O", 50)
	err = workbook.SetColWidth(sheetName, "P", "P", 50)
	err = workbook.SetRowHeight(sheetName, 1, 45)
	fmt.Println("Sizing set...")

	// set sheet base styling
	borders := []excelize.Border{{Type: "top", Color: "#000000", Style: 1}, {Type: "left", Color: "#000000", Style: 1}, {Type: "right", Color: "#000000", Style: 1}, {Type: "bottom", Color: "#000000", Style: 1}}
	style, err := workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 11, Family: "Calibri"}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}, Border: borders})
	err = workbook.SetColStyle(sheetName, "A:K", style)
	style, err = workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 11, Bold: true, Color: "#FFFFFF", Family: "Calibri"}, Fill: excelize.Fill{Type: "pattern", Color: []string{"#00008B"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}, Border: borders})
	err = workbook.SetCellStyle(sheetName, "A1", "K1", style)
	fmt.Println("Styles set...")

	// insert headers
	headers := []string{"Contract Start Month", "State", "Utility", "Zone", "Rate Code(s)", "Product Special Options", "Billing Method", "Term", "0 - 49", "50 - 299", "300 - 1099"}
	err = workbook.SetSheetRow(sheetName, "A1", &headers)
	fmt.Println("Headers inserted...")

	style, err = workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 11, Family: "Calibri"}, Fill: excelize.Fill{Type: "pattern", Color: []string{"#eeeeee"}, Pattern: 1}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}, Border: borders})

	// insert pricing entries
	rowIndex := 2
	for _, entry := range entries {
		err = workbook.SetRowHeight(sheetName, rowIndex, 90)
		startCell := "A" + strconv.Itoa(rowIndex)
		entrySlice := []string{entry.ContractStart, entry.State, entry.Util, entry.Zone, entry.RateCode, entry.ProductOption, entry.BillingMethod, strconv.Itoa(entry.Term), fmt.Sprintf("%.5f", entry.UsageLower), fmt.Sprintf("%.5f", entry.UsageMiddle), fmt.Sprintf("%.5f", entry.UsageUpper)}
		// alternate colors for easier reading
		if rowIndex%2 == 1 {
			err = workbook.SetRowStyle(sheetName, rowIndex, rowIndex, style)
		}
		err = workbook.SetSheetRow(sheetName, startCell, &entrySlice)
		fmt.Println("Entry inserted: " + startCell)
		rowIndex++
	}

	// get date and create text declaring Utility, Start Month/Year, and date of matrix pricing
	date, err := workbook.GetCellValue(mainSheetName, "A3")
	date = strings.ReplaceAll(date, "as of ", "")
	params := dbModify.ReadJson()
	var startDate string
	if params.StartDate != "" {
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
	} else {
		startDate = "ALL START"
	}
	// insert text
	infoText := fmt.Sprintf("%s %s Start (%s)", params.Util, startDate, date)
	infoStartCell := "A" + strconv.Itoa(rowIndex)
	infoEndCell := "K" + strconv.Itoa(rowIndex+3)
	err = workbook.MergeCell(sheetName, infoStartCell, infoEndCell)
	style, err = workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 24, Bold: true}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}, Border: borders})
	err = workbook.SetCellStyle(sheetName, infoStartCell, infoStartCell, style)
	err = workbook.SetSheetRow(sheetName, infoStartCell, &[]interface{}{infoText})

	//insert util name
	utilStartCell := "A" + strconv.Itoa(rowIndex+4)
	utilEndCell := "K" + strconv.Itoa(rowIndex+6)
	err = workbook.MergeCell(sheetName, utilStartCell, utilEndCell)
	style, err = workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 20, Bold: true}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}, Border: borders})
	utilName := dbModify.GetUtilByCode(params.Util)
	err = workbook.SetCellStyle(sheetName, utilStartCell, utilStartCell, style)
	err = workbook.SetSheetRow(sheetName, utilStartCell, &[]interface{}{utilName})

	fmt.Println("Info text inserted...")

	// clean up styling
	style, err = workbook.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 11, Family: "Calibri"}, Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true}})
	err = workbook.SetRowStyle(sheetName, rowIndex+7, rowIndex+54, style)
	err = workbook.SetColStyle(sheetName, "L:P", style)

	// save file
	err = workbook.Save()
	if err != nil {
		fmt.Printf("ERROR: %s! See matrixTool generate --help for more information", err)
		os.Exit(1)
	}
	fmt.Println("File Saved at: " + filePath)

}
