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
	fmt.Println("asap")
	dbModify.ProcessRow(rows[53])
	//for _, row := range rows[55:] {
	//	modifier.ProcessRow(row)
	//}
}
