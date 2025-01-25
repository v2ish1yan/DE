package tools

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

// 遍历读取每个单元格的数据
func Get_xlsx_cell(file string) []string {

	var cellvalues []string
	f, err := excelize.OpenFile(file)
	if err != nil {
		panic(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheets := f.GetSheetList()

	for _, sheet := range sheets {

		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Printf("读取工作表 %s 时出错: %v", sheet, err)
			continue
		}
		for _, row := range rows {
			for _, cell := range row {
				cellvalues = append(cellvalues, cell)
			}
		}
	}
	cellvalues = removeEmptyStrings(cellvalues)
	return cellvalues
}
