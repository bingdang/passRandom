# passRandom

>实现一个密码生成工具，支持以下功能：
用户可以通过 -l 指定生成密码的长度
用户可以通过 -t 指定生成密码的字符集，例如 -t num 生成全是数字的密码，-t char生成全是英文字符的密码，-t mix 生成包含数字和英文字符的密码，-t advance生成包含数字、英文以及特殊字符的密码
提示：用标准包flag解析命令行参数实现，用于生成密码的字符串如下：

demo
```
go build main.go

felix@MacBook-Pro project02 % ./main -h       
Usage of ./main:
  -l int
        生成密码的长度 (default 8)
  -t string
        advance: 字母+数字+特殊字符｜mix: 字母+数字｜char: 数字 (default "advance")
felix@MacBook-Pro project02 % ./main -t mix -l 12
sBsJKljlbSIi
felix@MacBook-Pro project02 % ./main -t char -l 12
763864239177
felix@MacBook-Pro project02 % ./main -l 12        
D!7,GRCsVv9U
felix@MacBook-Pro project02 % ./main      
#x06Sgrs

```
