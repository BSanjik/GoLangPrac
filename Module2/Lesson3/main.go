package main

import (
	"fmt"
	"strconv"
	"strings"
)

//add - returns sum of 2 numbers
// func add(a, b int) int {
// 	return a + b
// }

//variadic functions
// func sum(numbers ...int) int {
// 	var sum int
// 	for _, v := range numbers {
// 		sum += v
// 	}
// 	return sum
// }

//	func exampleAny(anySlice ...any) {
//		for _, v := range anySlice {
//			fmt.Println(v)
//		}
//	}\
//
// именнованный возврат данных
func add(a, b int) (result int) {
	result = a + b
	return
}

// Variatic
func summ(text string, numbers ...int) string {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return text + strconv.Itoa(sum)
}

func main() {
	// sliceInts := []int{6, 7, 8}
	// sum(sliceInts...)
	// exampleAny()
	// defer fmt.Println("1") // исполнится в конце
	// fmt.Println("2")
	// fmt.Println("3")

	// dbConn, err := sql.Open("postgres", "login:password:port:url") //data base connection
	// defer close(dbConn)
	//task1()
	// task2()
	// task3()
	// task4()
	// task5()
	// task6()
	// lessonTask1()
	// lessonTask2()
	// lessonTask3()

	// fmt.Println(add(1, 5))
	fmt.Print(summ("text ", 1, 2))
}

func lessonTask1() {
	var count int
	var temp, divider string
	fmt.Scan(&count)
	slice := make([]string, count)
	for i := range count {
		fmt.Scan(&temp)
		slice[i] = temp
	}
	fmt.Scan(&divider)
	fmt.Println(lessonTaskFunc(divider, slice...))
}
func lessonTaskFunc(divider string, words ...string) string {
	var result string
	result = strings.Join(words, divider)
	return result
}

func lessonTask2() {
	var temp int
	slice := make([]int, 5)
	for i := range slice {
		fmt.Scan(&temp)
		slice[i] = temp
	}
	fmt.Println(lessonTask2Func(slice...))
}

func lessonTask2Func(numbers ...int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] <= numbers[i-1] {
			return false
		}
	}
	return true
}

func lessonTask3() {
	var temp int
	slice := make([]int, 6)
	for i := range slice {
		fmt.Scan(&temp)
		slice[i] = temp
	}
	fmt.Println(lessonTask3Func(slice...))
}

func lessonTask3Func(numbers ...int) []int {
	var result []int
	var temp int
	for i := range numbers {
		temp = numbers[i]
		if temp == numbers[i] {
			result = append(result, numbers[i])
		}
	}
	return result
}

func task1() {
	var count, temp int
	fmt.Scan(&count)
	slice := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Scan(&temp)
		slice[i] = temp
	}

	fmt.Println(task1Summ(slice...))
}

func task1Summ(numbers ...int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func task2() {
	var count, nums int
	fmt.Scan(&count)
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&nums)
		slice[i] = nums
	}
	fmt.Println(task2Max(slice...))
}

func task2Max(numbers ...int) int {
	max := 0
	for i := range numbers {
		if max < i {
			max = numbers[i]
		}
	}
	return max
}

func task3() {
	var count, temp int
	fmt.Scan(&count)
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&temp)
		slice[i] = temp
	}
	fmt.Println(task3Sum(slice...))
}

func task3Sum(numbers ...int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func task4() {
	var count, temp int
	fmt.Scan(&count)
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&temp)
		slice[i] = temp
	}
	fmt.Println(task4Even(slice...))
}

func task4Even(numbers ...int) bool {
	var a bool
	for i := range numbers {
		if i%2 == 0 {
			a = true
			break
		}
	}
	return a
}

func task5() {
	var count, temp int
	fmt.Scan(&count)
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&temp)
		slice[i] = temp
	}
	fmt.Println(task5Multi(slice...))
}

func task5Multi(numbers ...int) int {
	result := 1
	for i := range numbers {
		result *= numbers[i]
	}
	return result
}

func task6() {
	var count, temp int
	fmt.Scan(&count)
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&temp)
		slice[i] = temp
	}
	fmt.Println(task6NotPos(slice...))
}

func task6NotPos(numbers ...int) bool {
	var result bool
	for i := range numbers {
		if numbers[i] < 0 {
			result = true
			break
		}
		result = false
	}
	return result
}
