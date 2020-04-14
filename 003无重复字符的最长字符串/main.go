package main

import (
	"fmt"
	"strings"
)

func main(){
	var a string = "abcdct"
	res := lengthOfLongestSubstring(a)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int{
	num, j, t := 0,0,0
	for i:=0;i<len(s);i++{
		//fmt.Println(s[j:i], string(s[i]))
		t = strings.Index(s[j:i], string(s[i]))
		//fmt.Println(t)
		if t == -1{
			if num < (i-j+1){
				num=i-j+1
				//fmt.Println(num)
				//fmt.Println()
			}
		}else {
			//fmt.Println(j)
			j=j+t+1
			//fmt.Println(j)
		}
	}
	return num
}
