package tools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Data string
}

func Output(data []map[string][]string, out string) {
	// 输出到csv

	resultdata := MergeMaps(data)
	if strings.HasSuffix(out, ".txt") {
		WriteRecordsToTXT(out, resultdata)
	} else {
		Write_csv_cell(out, resultdata)
	}
	PrintinConsole(resultdata)

}

// 使用github.com/gocarina/gocsv写入csv文件
func Write_csv_cell(filepath string, data map[string][]string) {

	csvdata := WriteforCsv(data)
	csvfile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	// 确保文件在函数结束时关闭
	writer := csv.NewWriter(csvfile)
	defer writer.Flush()
	for _, record := range csvdata {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return
		}
	}
	fmt.Println("CSV 文件写入成功！")
}

func WriteforCsv(data map[string][]string) [][]string {
	var Csvdata [][]string
	//计算最大的列数
	maxlengt := 0

	for _, record := range data {
		if len(record) > maxlengt {
			maxlengt = len(record)
		}
	}
	// 添加表头
	headers := []string{"电话号码", "身份证号码", "URL链接", "IP地址"}
	Csvdata = append(Csvdata, headers)

	//四列 "Phone", "ID Card", "URL", "IP"
	for i := 0; i < maxlengt; i++ {
		row := make([]string, 4)
		if i < len(data["phone"]) {
			row[0] = data["phone"][i]

		}
		if i < len(data["id_card"]) {
			row[1] = data["id_card"][i]
		}
		if i < len(data["url"]) {
			row[2] = data["url"][i]
		}
		if i < len(data["ip"]) {
			row[3] = data["ip"][i]
		}
		Csvdata = append(Csvdata, row)
	}

	return Csvdata
}

func WriteRecordsToTXT(filepath string, data map[string][]string) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("创建文件失败: %v\n", err)
		return
	}
	defer file.Close()
	// 写入数据到文件
	for k, v := range data {
		if k == "phone" {
			file.WriteString("=== 电话号码 ===\n")
			for _, phone := range v {
				_, err = file.WriteString(phone + "\n")
				if err != nil {
					panic(err)
				}
			}

		}
		if k == "id_card" {
			file.WriteString("=== 身份证 ===\n")
			for _, id_card := range v {
				_, err = file.WriteString(id_card + "\n")
				if err != nil {
					panic(err)
				}
			}
		}
		if k == "url" {
			file.WriteString("=== URL链接 ===\n")
			for _, url := range v {
				_, err = file.WriteString(url + "\n")
				if err != nil {
					panic(err)
				}
			}
		}
	}

	fmt.Printf("数据已成功写入%s\n", filepath)
}

func PrintinConsole(data map[string][]string) {
	fmt.Println("结果如下:")
	for k, v := range data {
		if k == "phone" {
			fmt.Println("=== 电话号码 ===\n")
			for _, phone := range v {
				fmt.Println(phone + "\n")

			}

		}
		if k == "id_card" {
			fmt.Println("=== 身份证 ===\n")
			for _, id_card := range v {
				fmt.Println(id_card + "\n")

			}
			if k == "url" {
				fmt.Println("=== URL链接 ===\n")
				for _, url := range v {
					fmt.Println(url + "\n")
				}
			}
		}
	}
}

func MergeMaps(maps []map[string][]string) map[string][]string {
	mergedMap := make(map[string][]string)
	for _, m := range maps {
		for k, v := range m {
			// 如果键已存在，合并切片
			if existingVal, exists := mergedMap[k]; exists {
				mergedMap[k] = append(existingVal, v...)
			} else {
				mergedMap[k] = v
			}
		}
	}
	return mergedMap
}
