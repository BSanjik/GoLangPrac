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
	case -2:
		// global scope
		sayHello()
	case -1:
		var weight float64
		var height float64
		fmt.Scan(&weight, &height)
		fmt.Println(classifyBMI(weight, height))
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

//calculateBMI - calculating BMI
func calculateBMI(weight float64, height float64) float64 {
	heightM := height / 100
	bmi := weight / (heightM * heightM)
	return bmi
}

//classifyBMI - classify BMI
func classifyBMI(weight float64, height float64) (string, float64) {

	result := calculateBMI(weight, height)
	var resultStr string
	if result < 18.5 {
		resultStr = "Недовес"
	} else if result >= 18.5 && result <= 24.9 {
		resultStr = "Норма"
	} else if result >= 25.0 && result <= 29.9 {
		resultStr = "Избыточный вес"
	} else {
		resultStr = "Жирный"
	}
	return resultStr, result
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
