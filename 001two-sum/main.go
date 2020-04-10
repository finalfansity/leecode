package main

import "fmt"

/*
给定一个整数数组nums和一个目标target，请在该数组中找出和为目标值的那两个整数，
并且返回他们的数组，可以假设每种输入只会对应一个答案， 但是不能重复利用这个数
组中同样的元素。

nums = [2,7,11,15]  target = 9
nums[0] + nums[1] = 2 + 7 = 9
所以返回[0,1]
 */

func twoNumSum(nums []int, target int)[]int{
	for i := 0; i < len(nums); i++ {
		//fmt.Println(i)
		for l:= i+1; l < len(nums); l++ {
			//fmt.Println(l)
			if nums[i] + nums[l] == target {
				return []int{i, l}
			}
		}
	}
	return []int{}
}

func main(){
	nums := []int{2, 7, 11, 15}
	target_num := 9

	res := twoNumSum(nums,target_num)
	fmt.Println(res)
}