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
	lastRowIdx := 1
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
			//将文件第一列 电影名字 缓存，用于判断文件中是否重复
			colData, _ := f.GetCols(sheetName)
			if colData != nil {
				movieNames := colData[0]
				lastRowIdx = len(movieNames) //历史数据最后一行id
				for _, movieName := range movieNames {
					existMovies[movieName] = 1
				}
			}
		}
		//设置默认工作表
		f.SetActiveSheet(index)
	}

	//set value of a cell
	newRowIdx := 0
	for _, row_data := range data {
		for col_idx, col_val := range row_data {
			if col_idx == 0 {
				_, ok := existMovies[col_val]
				if ok {
					//电影名字已存在
					//fmt.Printf(col_val)
					break
				}
				//新增行号计数器
				newRowIdx++
			}
			write_line_id := newRowIdx + lastRowIdx
			//转换成 A1 B1
			col_key := string(65+col_idx) + strconv.Itoa(write_line_id)
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
