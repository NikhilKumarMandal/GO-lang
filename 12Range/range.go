package main

import (
	"fmt"
)

func main() {
	//nums := []int{5, 6, 7}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// i -> index
	// num -> value
	// for i,num := range nums {
	// 	sum += num
	// 	fmt.Println(i)
	// }

	// fmt.Println(sum)

	// m := map[string]string{"fname": "john","lname":"doe"}

	// for k,v := range m {
	// 	fmt.Println(k,v)
	// }

	// for _,v := range m {
	// 	fmt.Println(v)
	// }

	// unicode code point rune
	// starting byte of rune
	// 255  => 1 byte , if grater then 255 it will take more than 1 byte
	for i,c := range "golang"{
		fmt.Println(i,string(c))
	}
}