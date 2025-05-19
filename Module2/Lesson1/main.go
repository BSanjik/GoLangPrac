package main

import "fmt"

const (
	trueLogin = "admin"
	truePass  = "admin123"
)

func main() {

	fmt.Println("Write prac number")
	var pracNum int
	fmt.Scan(&pracNum)
	switch pracNum {
	case 0:
		var login string
		var pass string
		fmt.Scan(&login, &pass)
		if isAuth(login, pass) == false {
			fmt.Println("BAD")
		} else {
			fmt.Println("YES")
		}
	case 1:
		var n int
		var m int
		fmt.Scan(&n)
		fmt.Scan(&m)
		fmt.Println(prac1(n, m))
	case 2:
		var n string
		fmt.Scan(&n)
		fmt.Println(prac2(n))
	case 3:
		var n int
		fmt.Scan(&n)
		fmt.Println(prac3(n))
	}
}

func isAuth(login string, pass string) bool {

	if login == trueLogin && pass == truePass {
		return true
	} else {
		return false
	}

}

func prac1(a, b int) int {
	return a + b
}

func prac2(n string) int {
	lenString := []rune(n)
	return len(lenString)
}

func prac3(n int) int {
	return n * n
}
