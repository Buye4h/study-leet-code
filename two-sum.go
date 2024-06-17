package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	//create new map for keep key of num
	numMap := make(map[int]int)

	//loop nums
	// i = index ของเลขปัจจุบันที่กำลังเข้า loop
	// num = ค่าของเลขนั้น
	for i, num := range nums {
		//เอาเลขโจทย์มาลบเพื่อหา "ส่วนต่าง"
		complement := target - num
		//เช็คว่า "ค่าส่วนต่าง" นั้น มีอยู่ใน map ป่าว

		if k, ok := numMap[complement]; ok {
			// มี -> ส่ง index ของ k,i
			return []int{k, i}
		}
		// ไม่มี -> เก็บว่าเลขนั้นมีค่า index ตรงไหน -> วนลูปต่อไป
		numMap[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{3, 2, 5, 3}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
