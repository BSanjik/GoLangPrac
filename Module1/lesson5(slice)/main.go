package main

import "fmt"

func main() {
	//n1
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	//	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range nums {
	// 		fmt.Scan(&nums[i])
	// 	}
	// 	fmt.Println(nums[3])
	// }

	//n2
	// var n, max int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	max = nums[0]
	// 	for i := range nums {
	// 		fmt.Scan(&nums[i])
	// 		if max < nums[i] {
	// 			max = nums[i]
	// 		}
	// 	}
	// 	fmt.Println(max)
	// }

	//n3
	// var n, result int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	result = 1
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if nums[i] != 0 {
	// 			result *= nums[i]
	// 		}
	// 	}
	// }
	// fmt.Println(result)

	//n4
	// var n, min int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if i == 0 {
	// 			min = nums[i]
	// 		} else {
	// 			if min > nums[i] {
	// 				min = nums[i]
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(min * min)

	//n5
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	max := nums[0]
	// 	tempIndex := 0
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if max < nums[i] {
	// 			max = nums[i]
	// 			tempIndex = i
	// 		}
	// 	}
	// 	fmt.Println(max, tempIndex)
	// }

	//n6
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if i == 0 {
	// 			min, max = nums[i], nums[i]
	// 		} else {
	// 			if min > nums[i] {
	// 				min = nums[i]
	// 			}
	// 			if max < nums[i] {
	// 				max = nums[i]
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(min, max)
	// }

	//n7
	// var n, min, max int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if i == 0 {
	// 			min, max = nums[i], nums[i]
	// 		} else {
	// 			if min > nums[i] {
	// 				min = nums[i]
	// 			}
	// 			if max < nums[i] {
	// 				max = nums[i]
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(max - min)
	// }

	//n8
	// var n, k int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 	}
	// 	fmt.Scan(&k)
	// 	for i := range n {
	// 		if nums[i]%k == 0 {
	// 			fmt.Println(nums[i])
	// 		}
	// 	}
	// }

	//n9
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 	}
	// 	for i := n - 1; i >= 0; i-- {
	// 		fmt.Println(nums[i])
	// 	}
	// }

	//n10
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	numsSec := []int{}
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if nums[i]%2 == 0 {
	// 			numsSec = append(numsSec, nums[i])
	// 		}
	// 	}
	// 	fmt.Println(numsSec)
	// }

	//n11
	// var n, result int
	// _, err := fmt.Scan(&n)
	// if err != nil{
	// 	fmt.Println("Error")
	// 	return
	// }else{
	// 	nums := make([]int, n)
	// 	for i := range n{
	// 		fmt.Scan(&nums[i])
	// 		result += nums[i]
	// 	}
	// 	fmt.Println(result)
	// }

	//n12
	// var n int
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if i == 0 && nums[i] < 0 {
	// 			fmt.Println("Error")
	// 			return
	// 		} else {
	// 			if i%2 != 0 && nums[i] > 0 {
	// 				fmt.Println("NO")
	// 				return
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("YES")
	// }

	//n13
	// initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	// users := initialUsers[0:4]

	// for i := range len(users) {
	// 	fmt.Print(users[i], " ")
	// }

	//n14
	// var n int
	// result := 1
	// _, err := fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	nums := make([]int, n)
	// 	for i := range n {
	// 		fmt.Scan(&nums[i])
	// 		if i%2 == 0 {
	// 			result *= nums[i]
	// 		}
	// 	}
	// 	fmt.Println(result)
	// }

	//n15
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error")
		return
	} else {
		names := make([]string, n)
		for i := range n {
			fmt.Scan(&names[i])
		}
		names = append(names, "Arman", "Ilyas", "Tore")
		for i := len(names) - 1; i > 0; i-- {
			fmt.Println(names[i])
		}
	}
}

// func LessonTask1() {
// 	nums_1 := []int{0, 1, 2, 3}
// 	nums_2 := []int{4, 5, 6, 7}

// 	nums_1 = append(nums_1, nums_2...)

// 	for i := range len(nums_1) {
// 		fmt.Print(nums_1[i], " ")
// 	}
// }

// func LessonTask2() {
// 	nums_1 := []int{0, 1, 2, 3, 4}
// 	var user_input int

// 	fmt.Println("Введите индекс числа которое нужно удалить")

// 	fmt.Fscan(os.Stdin, &user_input)

// 	output_nums := append(nums_1[0:user_input], nums_1[user_input+1:]...)

// 	for i := range len(output_nums) {
// 		fmt.Print(output_nums[i], " ")
// 	}
// }
// }
