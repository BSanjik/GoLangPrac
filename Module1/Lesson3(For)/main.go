package main

import "fmt"

func main() {
	//n1
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("I like BITLAB!!!")
	// }

	//n2
	// var n int
	// fmt.Scan(&n)
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d ", n)
	// }

	//n3
	// var n int
	// fmt.Scan(&n)
	// for i := 1; i < n+1; i++ {
	// 	fmt.Printf("%d ", i)
	// }

	//n4
	// var n int
	// var str string
	// fmt.Scan(&str, &n)
	// for i := 1; i < n+1; i++ {
	// 	fmt.Println(str)
	// }

	//n5
	// var n, m int
	// fmt.Scan(&n, &m)
	// for i := n; i < m+1; i++ {
	// 	fmt.Printf("%d ", i)
	// }

	//n6
	// // for i := 1; i <= 10; i++ {
	// // 	fmt.Println(i * i)
	// }

	//n7
	// var n int
	// fmt.Scan(&n)
	// for i := 1; i <= n; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//n8
	// var n int
	// fmt.Scan(&n)
	// result := 1
	// for i := 1; i < n+1; i++ {
	// 	result = result * i
	// }
	// fmt.Println(result)

	//n9
	// var a, b int
	// result := 0
	// _, err := fmt.Scan(&a, &b)
	// if err != nil || a > 101 && b > 101 && a < b {
	// 	fmt.Println("Error")
	// 	return
	// }
	// for i := a; i < b+1; i++ {
	// 	result = result + i
	// }
	// fmt.Print(result)

	//n10
	// var n, c, d int
	// _, err := fmt.Scan(&n, &c, &d)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }
	// for i := 1; i <= n+1; i++ {
	// 	if i%c == 0 && i%d != 0 {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// }

	//n11
	// var n, m int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }
	// for i := n; i < m+1; i++ {
	// 	if i%2 > 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//n12
	// var n, m int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }
	// for i := n; i < m+1; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//n13
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }
	// for i := 1; i <= n; i++ {
	// 	if n%i == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//n14
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }
	// for i := 1; i <= n; i++ {
	// 	fmt.Println(i, " ", i*i)
	// }

	//n15
	var a, b, c, d int
	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		fmt.Println("Error")
		return
	}
	for i := a; i <= b; i++ {
		if i%d == c {
			fmt.Println(i)
		}
	}

}
