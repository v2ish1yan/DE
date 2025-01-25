package main

import (
	"flag"
	"fmt"
	"foriensec/common"
	"os"
)

func main() {
	//接收命令行参数
	path := flag.String("path", "", "文件名或目录名")
	output := flag.String("out", "output.csv", "输出文件格式：csv、txt")
	flag.Usage = func() {
		// 显示艺术字
		art := `
          _____                    _____          
         /\    \                  /\    \         
        /::\    \                /::\    \        
       /::::\    \              /::::\    \       
      /::::::\    \            /::::::\    \      
     /:::/\:::\    \          /:::/\:::\    \     
    /:::/  \:::\    \        /:::/__\:::\    \    
   /:::/    \:::\    \      /::::\   \:::\    \   
  /:::/    / \:::\    \    /::::::\   \:::\    \  
 /:::/    /   \:::\ ___\  /:::/\:::\   \:::\    \ 
/:::/____/     \:::|    |/:::/__\:::\   \:::\____\
\:::\    \     /:::|____|\:::\   \:::\   \::/    /
 \:::\    \   /:::/    /  \:::\   \:::\   \/____/ 
  \:::\    \ /:::/    /    \:::\   \:::\    \     
   \:::\    /:::/    /      \:::\   \:::\____\    
    \:::\  /:::/    /        \:::\   \::/    /    
     \:::\/:::/    /          \:::\   \/____/     
      \::::::/    /            \:::\    \         
       \::::/    /              \:::\____\        
        \::/____/                \::/    /        
         ~~                       \/____/         
                                                  
`
		_, _ = fmt.Fprintln(os.Stderr, art)
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		_, _ = fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		_, _ = fmt.Fprintln(os.Stderr, "Example:main.exe -path target.csv -out output.csv")
		_, _ = fmt.Fprintln(os.Stderr, "")
	}
	////解析命令行参数
	flag.Parse()
	//path := "tools/output.csv"
	//xlsxdata := tools2.Get_csv_cell(path)
	//reg_res := tools2.Reg_all(xlsxdata)
	//tools2.Output(reg_res)
	common.Start(*path, *output)

}
