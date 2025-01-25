package tools

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// 使用encoding/csv 读取csv文件

func Get_csv_cell(filepath string) []string {

	var cellvalues []string
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建一个新的 CSV 读取器
	reader := csv.NewReader(file)

	// 读取 CSV 文件中的所有记录
	for {
		// 读取一行记录
		record, err := reader.Read()
		cellvalues = append(cellvalues, record...)
		if err == io.EOF {
			break // 到达文件末尾
		}
		if err != nil {
			log.Fatalf("读取记录时出错: %v", err)
		}

		// 打印记录
		//fmt.Println(record)

	}
	return cellvalues
}
