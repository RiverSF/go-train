package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

func UpdateExcel(data [][]string, path string, sheetName string) {
	if data == nil || path == "" {
		fmt.Println()
		return
	}

	//判断文件是否存在
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		CreateExcel(data, path, sheetName)
		return
	}

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	//create a new sheet
	existMovies := make(map[string]int)
	beginRow := 1
	if sheetName == "" {
		sheetName = "Sheet1"
	} else {
		//判断sheet name 是否重复
		index, _ := f.GetSheetIndex(sheetName)
		if index == -1 {
			index, err = f.NewSheet(sheetName)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			colData, _ := f.GetCols(sheetName)
			if colData != nil {
				movieNames := colData[0]
				beginRow = len(movieNames) + 1
				for _, movieName := range movieNames {
					existMovies[movieName] = 1
				}
			}
		}
		//设置默认工作表
		f.SetActiveSheet(index)
	}

	//set value of a cell
	for row_id, row_data := range data {
		write_line_id := row_id + beginRow
		for col_id, col_val := range row_data {
			if col_id == 0 {
				_, ok := existMovies[col_val]
				if ok {
					//电影名字已存在
					//fmt.Printf(col_val)
					break
				}
			}
			//转换成 A1 B1
			col_key := string(65+col_id) + strconv.Itoa(write_line_id)
			//fmt.Println(col_idx, cell_val, col_key)
			err := f.SetCellValue(sheetName, col_key, col_val)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	//save spreadsheet by the given path
	//excel/excel.xlsx
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return
	}

}
