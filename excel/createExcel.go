package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func CreateExcel(data [][]string, path string, sheetName string) {
	if data == nil || path == "" {
		fmt.Println()
		return
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	//create a new sheet
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		index, err := f.NewSheet(sheetName)
		if err != nil {
			fmt.Println(err)
			return
		}
		//设置默认工作表
		f.SetActiveSheet(index)
	}

	// mock data
	//data := make([][]string, 10)
	//for i := 0; i < 10; i++ {
	//	row := make([]string, 5)
	//	for j := 0; j < 5; j++ {
	//		if i == 0 {
	//			row[j] = "表头" + strconv.Itoa(j)
	//			continue
	//		}
	//		row[j] = strconv.Itoa(i * j)
	//	}
	//	data[i] = row
	//}

	//set value of a cell
	for line_idx, rows := range data {
		first_line_idx := line_idx + 1
		//fmt.Println(first_line_idx)
		for col_idx, cell_val := range rows {
			col_key := string(65+col_idx) + strconv.Itoa(first_line_idx)
			//fmt.Println(col_idx, cell_val, col_key)
			err := f.SetCellValue(sheetName, col_key, cell_val)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	//save spreadsheet by the given path
	//excel/excel.xlsx
	if err := f.SaveAs(path); err != nil {
		fmt.Println(err)
	}

}
