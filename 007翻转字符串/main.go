package main

import "fmt"

/*
问题描述
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
*/
func main(){
   s := "测试中数据"
   res, f := reverStr(s)
   fmt.Println(res,f)
}


/*
  1   2   3   4   5
  3 为中
  1,5  -> 5,1
*/
func reverStr(s string)(string, bool){
	//一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
	//另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。
	str := []rune(s)
	l := len(str)
	if l > 5000{
		return s, false
	}
	for i:=0; i<l/2;i++{
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str), true
}
