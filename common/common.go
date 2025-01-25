package common

import (
	"fmt"
	"foriensec/tools"
	"os"
	"path/filepath"
	"strings"
)

func Start(path string, output string) {

	var results []map[string][]string
	//默认输出格式为csv
	if output == "" {
		output = "output.csv"
	}

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("文件或目录不存在")
		return
	}

	ext1 := strings.ToLower(filepath.Ext(path))
	switch ext1 {
	case ".xlsx":
		fmt.Println("开始处理文件：", path)
		xlsxdata := tools.Get_xlsx_cell(path)
		reg_res := tools.Reg_all(xlsxdata)
		results = append(results, reg_res)

	case ".txt":
		fmt.Println("开始处理文件：", path)
		txtdata := tools.Get_txt_data(path)
		reg_res := tools.Reg_all(txtdata)
		results = append(results, reg_res)

	case ".csv":
		fmt.Println("开始处理文件：", path)
		csvdata := tools.Get_csv_cell(path)
		reg_res := tools.Reg_all(csvdata)
		results = append(results, reg_res)

	default:
		{
			fmt.Println("文件格式不支持:", path)
		}
	}
	if info.IsDir() {
		fmt.Println("开始遍历目录文件")
		err = filepath.Walk(path, func(file_list string, info os.FileInfo, err error) error {
			if file_list != path && !info.IsDir() {
				ext := strings.ToLower(filepath.Ext(file_list))
				switch ext {
				case ".xlsx":
					fmt.Println("开始处理文件：", file_list)
					xlsxdata := tools.Get_xlsx_cell(file_list)
					reg_res := tools.Reg_all(xlsxdata)
					results = append(results, reg_res)

				case ".txt":
					fmt.Println("开始处理文件：", file_list)
					txtdata := tools.Get_txt_data(file_list)
					reg_res := tools.Reg_all(txtdata)
					results = append(results, reg_res)

				case ".csv":
					fmt.Println("开始处理文件：", file_list)
					csvdata := tools.Get_csv_cell(file_list)
					reg_res := tools.Reg_all(csvdata)
					results = append(results, reg_res)

				default:
					{
						fmt.Println("文件格式不支持:", file_list)
					}
				}

			}
			return nil
		})

	}
	tools.Output(results, output)

	//fmt.Println("start")
}
