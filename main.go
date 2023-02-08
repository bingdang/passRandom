package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

//实现一个密码生成工具，支持以下功能：
//用户可以通过 -l 指定生成密码的长度
//用户可以通过 -t 指定生成密码的字符集，例如 -t num 生成全是数字的密码，-t char生成全是英文字符的密码，-t mix 生成包含数字和英文字符的密码，-t advance生成包含数字、英文以及特殊字符的密码
//提示：用标准包flag解析命令行参数实现，用于生成密码的字符串如下：

var (
	Len int
	Chr string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

//读取终端参数
func Init() {
	rand.Seed(time.Now().Unix())
	flag.IntVar(&Len, "l", 8, "生成密码的长度")
	flag.StringVar(&Chr, "t", "advance", "advance: 字母+数字+特殊字符｜mix: 字母+数字｜char: 数字")
	flag.Parse()
}

//将字符串转为切片
func encapsulation(str string) (slice []byte) {
	for i := 0; i < len(str); i++ {
		slice = append(slice, str[i])
	}
	return slice
}

func structure() string {
	num := encapsulation(NumStr)
	cha := encapsulation(CharStr)
	spe := encapsulation(SpecStr)
	var passwd []byte = make([]byte, Len, Len)

	if Chr == "char" {
		num = num
	} else if Chr == "mix" {
		num = append(num, cha...)
	} else if Chr == "advance" {
		num = append(num, cha...)
		num = append(num, spe...)
	}
  
  //生成指定长度随机字符
	for i := 0; i < Len; i++ {
		index := rand.Intn(len(num))
		passwd[i] = num[index]
	}
	return string(passwd)
}

func main() {
	//初始化程序+随机数作种
	Init()

	//生成密码
	Passwd := structure()

	//输出密码
	fmt.Println(Passwd)
}
