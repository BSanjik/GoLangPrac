package main

import "fmt"

// func PrintInt(x int) {
// 	fmt.Println("x:", x)
// }

// func PrintAnything(x any) {
// 	fmt.Println("x:", x)
// }

func main() {
	// // var num int = 15
	// // var text string = "text"

	// // PrintAnything(num)
	// // PrintAnything(text)

	// var request any = 10

	// //type assertion
	// // result, ok := request.(string)
	// // if ok {
	// // 	fmt.Printf("%T - %v", result, result)
	// // 	return
	// // }
	// // fmt.Println("Error")

	// //switch - type
	// switch v := request.(type) {
	// case string:
	// 	fmt.Println("string", v)
	// case int:
	// 	fmt.Println("int", v)
	// default:
	// 	fmt.Println("Error")
	// }
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(Revers(&n, &m))
	fmt.Println(n, m)
}

// func Summ(num *int) int {
// 	*num += 1
// 	return *num
// }

func Revers(a, b *int) (int, int) {
	*a, *b = *b, *a
	return *a, *b
}
