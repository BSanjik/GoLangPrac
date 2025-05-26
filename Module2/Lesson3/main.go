package main

import "fmt"

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

// func exampleAny(anySlice ...any) {
// 	for _, v := range anySlice {
// 		fmt.Println(v)
// 	}
// }

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
	task5()

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
