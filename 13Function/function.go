package main

// func add(a,b int) int {
// 	return a + b
// }

// func genLanguages() (string,string,bool) {
// 	return "golang", "javascript", true
// }

// func processIt(fn func(a int) int){
// 	fn(3)
// }

func processIt() func (a int) int  {
	return func(a int) int{
		return 1
	}
}


func main() {
	// result := add(3, 5)
	// fmt.Println(result)

	// lang1,lang2,_ := genLanguages()
	// fmt.Println(lang1,lang2)
	// fn := func (a int) int  {
	// 	return 2
	// }
	fn := processIt()
	fn(2)
}
