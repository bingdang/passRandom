# passRandom
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
