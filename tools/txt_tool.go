package tools

import (
	"bufio"
	"fmt"
	"os"
)

func Get_txt_data(filepath string) []string {

	var lines []string
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("文件打开失败:%v\n", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件失败:%v\n", err)
		panic(err)
	}
	return lines
}
