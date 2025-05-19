package main

import "fmt"

func main() {
	//n1
	// var n int
	// //Error handler
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if n == 0 {
	// 	fmt.Println("Zero")
	// } else if n < 0 {
	// 	fmt.Println("Negative")
	// } else if n > 0 {
	// 	fmt.Println("Positive")
	// }

	//n2
	// var n int
	// fmt.Scan(&n)
	// if n > 999 || n < 100 {
	// 	fmt.Println("Число не трехзначное")
	// } else {
	// 	var num1 int = (n % 100) % 10
	// 	var num2 int = (n / 10) % 10
	// 	var num3 int = n / 100

	// 	if num1 != num2 && num1 != num3 && num2 != num3 {
	// 		fmt.Println("YES")
	// 	} else {
	// 		fmt.Println("NO")
	// 	}
	// }

	//n3
	// var n int
	// fmt.Scan(&n)
	// if n > 0 {
	// 	temp := strconv.Itoa(n)
	// 	fmt.Println(string(temp[0]))
	// }

	//n4
	// var n int
	// fmt.Scan(&n)

	// var num1, num2, num3, num4, num5, num6 int

	// num1 = n / 100000
	// num2 = (n / 10000) % 10
	// num3 = (n / 1000) % 10
	// num4 = (n / 100) % 10
	// num5 = (n / 10) % 10
	// num6 = n % 10

	// sum1 := num1 + num2 + num3
	// sum2 := num4 + num5 + num6

	// if sum1 == sum2 {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	//n5
	// var n int
	// fmt.Scan(&n)

	// condition1 := n % 400
	// condition2 := n % 4
	// condition3 := n % 100

	// if condition1 == 0 || (condition2 == 2 && condition3 != 0) {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	//n6
	// var n int
	// fmt.Scan(&n)

	// condition1 := n % 17

	// if condition1 == 0 {
	// 	fmt.Println("DIVISIBLE BY 17")
	// }

	//n7
	// var x, y int

	// fmt.Scan(&x, &y)

	// switch {
	// case x < y:
	// 	fmt.Println("<")
	// case x > y:
	// 	fmt.Println(">")
	// default:
	// 	fmt.Println("=")
	// }

	//n8
	// var x, y, result int

	// fmt.Scan(&x, &y, &result)

	// condition1 := x * y

	// if condition1 == result {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	//n9
	// var n int
	// fmt.Scan(&n)

	// if n > 12 || n <= 0 {
	// 	fmt.Println("Error")
	// 	return
	// }

	// switch n {
	// case 3, 4, 5:
	// 	fmt.Println("Spring")
	// case 6, 7, 8:
	// 	fmt.Println("Summer")
	// case 9, 10, 11:
	// 	fmt.Println("Autumn")
	// case 12, 1, 2:
	// 	fmt.Println("Winter")
	// }

	//n10
	// var x int
	// fmt.Scan(&x)

	// if x%2 == 0 {
	// 	fmt.Println("EVEN")
	// } else {
	// 	fmt.Println("ODD")
	// }

	//n11
	// var x, y int
	// fmt.Scan(&x, &y)

	// if x == y {
	// 	fmt.Println("0")
	// } else {
	// 	fmt.Println(x + y)
	// }

	//n12
	// var x int
	// fmt.Scan(&x)

	// if x > 11 || x < 1 {
	// 	fmt.Println("No such class")
	// 	return
	// }

	// switch x {
	// case 1, 2, 3, 4:
	// 	fmt.Println("Beginners School")
	// case 5, 6, 7, 8, 9:
	// 	fmt.Println("Middle School")
	// case 10, 11:
	// 	fmt.Println("High Shool")
	// }

	//n13
	// var x int
	// fmt.Scan(&x)

	// switch {
	// case x <= 2019:
	// 	fmt.Println("NO")
	// case x >= 2021:
	// 	fmt.Println("YES")
	// case x == 2020:
	// 	fmt.Println("ERROR")
	// }

	//n14
	// var x, y, z int
	// fmt.Scan(&x, &y, &z)

	// if x == y && x == z {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	//n15
	// var x int

	// fmt.Scan(&x)

	// if x > 0 && x%2 == 0 {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	//lesstask1
	// var a int
	// var b int
	// var c int

	// fmt.Scan(&a, &b, &c)

	// if a > b {
	// 	fmt.Println(b, a)
	// } else {
	// 	fmt.Println(a, b)
	// }

	//lesstask2
	// var Master_login string = "admin"
	// var Master_password string = "admin"

	// var Master_password_v2 string = "admin_2"

	// var login string
	// var password string

	// fmt.Println("Введите логин:")
	// fmt.Fscan(os.Stdin, &login)
	// fmt.Println("Введите пароль:")
	// fmt.Fscan(os.Stdin, &password)

	// if login == Master_login && (password == Master_password || password == Master_password_v2) {
	// 	fmt.Println("Успешно авторизовался!")
	// } else {
	// 	fmt.Println("Ну чет не правильно...")
	// }

	//lesstask3
	// var x int //1234
	// _, err := fmt.Scan(&x)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if x > 9999 || x < 999 {
	// 	fmt.Println("Число не подходит")
	// 	return
	// }

	// var num1, num2, num3, num4 int

	// num1 = x / 1000
	// num2 = (x / 100) % 10
	// num3 = (x / 10) % 10
	// num4 = x % 10

	// if num1+num2 == num3+num4 {
	// 	fmt.Println("LUCKY")
	// } else {
	// 	fmt.Println("TRY AGAIN")
	// }

	//lesstask4
	// var x int //123
	// _, err := fmt.Scan(&x)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if x > 999 || x < 100 {
	// 	fmt.Println("Число не подходит")
	// 	return
	// }

	// var num1, num3 int

	// num1 = x / 100
	// num3 = x % 10

	// if num1 == num3 {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	//lesstask5
	// var x, y int
	// fmt.Scan(&x, &y)

	// if x > 0 && y > 0 {
	// 	fmt.Println(x + y)
	// } else {
	// 	fmt.Println(x * y)
	// }

	//lesstask6
	var x int

	fmt.Scan(&x)

	if x > 99999 || x < 10000 {
		fmt.Println("Error. Incorrect Number")
		return
	}
	var num1, num2, num4, num5 int

	num1 = x / 10000
	num2 = (x / 1000) % 10
	num4 = (x / 10) % 10
	num5 = x % 10

	if num1 == num5 && num2 == num4 {
		fmt.Println("PALINDROME")
	} else {
		fmt.Println("NO")
	}
}
