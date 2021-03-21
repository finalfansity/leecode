package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。 给定一个string为原始的串，返回替换后的string。
*/
func main(){
	s1 := "wejfioae waje fioa jefoawef "
	res, f := replaceBlank(s1)
	fmt.Println(res, f)
}
func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000{
		return s, false
	}
	/*
	这里使用了golang内置方法unicode.IsLetter判断字符是否是字母，之后使用strings.Replace来替换空格。
	*/
	for _, v:=range s{
		if string(v) != " " && unicode.IsLetter(v) == false{
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
