package main

import (
	"fmt"
)

const (
	trueLogin = "admin"
	truePass  = "admin123"
)

func main() {

	fmt.Println("Write prac number")
	var pracNum int
	_, err := fmt.Scan(&pracNum)
	if err != nil {
		fmt.Println("Error")
		return
	} else {
		switch pracNum {
		case -2:
			// global scope
			//sayHello()
		case -1:
			var weight, height float64
			_, err := fmt.Scan(&weight, &height)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(classifyBMI(weight, height))
			}
		case 0:
			var login, pass string
			_, err := fmt.Scan(&login, &pass)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				if isAuth(login, pass) == false {
					fmt.Println("BAD")
				} else {
					fmt.Println("YES")
				}
			}
		case 1:
			var n, m int
			_, err := fmt.Scan(&n, &m)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac1(n, m))
			}
		case 2:
			var n string
			_, err := fmt.Scan(&n)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac2(n))
			}
		case 3:
			var n int
			_, err := fmt.Scan(&n)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac3(n))
			}
		case 4:
			var n, m int
			_, err := fmt.Scan(&n, &m)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac4(n, m))
			}
		case 5:
			var n int
			_, err := fmt.Scan(&n)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				nums := make([]int, n)
				for i := range n {
					fmt.Scan(&nums[i])
				}
				fmt.Println(prac5(nums))
			}
		case 6:
			var n int
			_, err := fmt.Scan(&n)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac6(n))
			}
		case 7:
			var text string
			_, err := fmt.Scan(&text)
			if err != nil {
				fmt.Println("Error")
				return
			}
			fmt.Println(prac7(text))
		case 8:
			strings := []string{"apple", "banana", "kiwi", "pineapple", "grape", "strawberry"}
			fmt.Println(prac8(strings))
		case 9:
			var n, m int
			_, err := fmt.Scan(&n, &m)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac9(n, m))
			}
		case 10:
			var n int
			_, err := fmt.Scan(&n)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac10(n))
			}
		case 11:
			nums := []int{1, 2, 2, 3, 4, 4, 5, 1}
			fmt.Println(prac11(nums))
		case 12:
			var n, m int
			_, err := fmt.Scan(&n, &m)
			if err != nil {
				fmt.Println("Error")
				return
			} else {
				fmt.Println(prac12(n, m))
			}
		case 13:
			text := "hello world"
			char := 'k'
			fmt.Println(prac13(text, char))
		case 14:
			nums := []int{-3, 0, 7, -1, 4, 9, -10}
			fmt.Println(prac14(nums))
		}
	}
}

// calculateBMI - calculating BMI
func calculateBMI(weight float64, height float64) float64 {
	heightM := height / 100
	bmi := weight / (heightM * heightM)
	return bmi
}

// classifyBMI - classify BMI
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

func prac4(a, b int) int {
	return a * b
}

func prac5(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	return result
}

func prac6(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func prac7(a string) string {
	var result string
	revers := []rune(a)
	for i := len(revers) - 1; i >= 0; i-- {
		result += string(revers[i])
	}
	return result
}

func prac8(arr []string) []string {
	var result []string
	for i := range arr {
		if i > 5 {
			result = append(result, arr[i])
		}
	}
	return result
}

func prac9(a, b int) int {
	result := a % b
	return result
}

func prac10(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func prac11(numsArr []int) []int {
	nums := make(map[int]bool)
	var result []int

	for i := range numsArr {
		if !nums[numsArr[i]] {
			nums[numsArr[i]] = true
			result = append(result, numsArr[i])
		}
	}
	return result
}

func prac12(a, b int) bool {
	if b == 0 {
		return false
	}
	return a%b == 0
}

func prac13(text string, char rune) int {
	count := 0
	for _, c := range text {
		if c == char {
			count++
		}
	}
	return count
}

func prac14(nums []int) []int {
	var result []int
	for i := range nums {
		if i > 0 {
			result = append(result, i)
		}
	}
	return result
}

// func prac15(text string) string {
// 	// neponyatno
// }
