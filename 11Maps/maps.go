package main

import (
	"fmt"
	"maps"
)

// mpas -> hash,object,dict

func main(){
	// creating map

	// m := make(map[string]string)

	// setting an element

	// m["name"] = "golang"
	// m["area"] = "backend"
	// get an element
	// fmt.Println(m["name"],m["area"])

	// IMP: IF key does not exists in the map it return zero value
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	//fmt.Println(m["phone"])

	// fmt.Println(len(m))

	// delete(m,"price")

	// clear((m))
	// fmt.Println(m)



	// fmt.Println(m)


	// m := map[string]int{"price":30,"phone":7878}

	// k,ok := m["price"]
	// fmt.Println(k)
	// if ok {
	// 	fmt.Println("all ok")
	// }else {
	// 	fmt.Println("not okay")
	// }

	m := map[string]int{"price":30,"phone":7878}
	m1 := map[string]int{"price":40,"phone":6969}

	fmt.Println(maps.Equal(m,m1))

}