## 功能

- 读取csv文件、xlsx文件、txt文件
- 正则匹配其中的手机号码、网址、身份证号
- 自动判断是否为文件，如果为目录则遍历目录下所有文件
- 输出结果到指定文件，支持csv、xlsx、txt格式

```bash

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


Usage of E:\code\GO\TOOLS\foriensec\main.exe:
Options:
  -out string
        输出文件格式：csv、txt (default "output.csv")
  -path string
        文件名或目录名
Example:main.exe -path target.csv -out output.csv

```
