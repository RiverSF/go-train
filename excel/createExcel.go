package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func CreateExcel() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//create a new sheet
	//index, err := f.NewSheet("Sheet2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// mock data
	data := make([][]string, 10)
	for i := 0; i < 10; i++ {
		row := make([]string, 5)
		for j := 0; j < 5; j++ {
			if i == 0 {
				row[j] = "表头" + strconv.Itoa(j)
				continue
			}
			row[j] = strconv.Itoa(i * j)
		}
		data[i] = row
	}
	//fmt.Println(data)
	//return

	//set value of a cell
	for line_idx, rows := range data {
		first_line_idx := line_idx + 1
		//fmt.Println(first_line_idx)
		for col_idx, cell_val := range rows {
			col_key := string(65+col_idx) + strconv.Itoa(first_line_idx)
			//fmt.Println(col_idx, cell_val, col_key)
			err := f.SetCellValue("Sheet1", col_key, cell_val)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	//设置默认工作表
	//f.SetActiveSheet(index)

	//save spreadsheet by the given path
	if err := f.SaveAs("excel/excel.xlsx"); err != nil {
		fmt.Println(err)
	}

}
