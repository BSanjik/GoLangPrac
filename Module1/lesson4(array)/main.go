package main

import "fmt"

func main() {

	//n1
	// var n int
	// nums := []int{6, 19, 26, 9, 46, 8, 5, 65, 90, 25}
	// _, err := fmt.Scan(&n)
	// if err != nil && n >= len(nums) || n < 0 {
	// 	fmt.Println("Error")
	// } else {
	// 	fmt.Println(nums[n])
	// }

	//n2
	// nums := []int{4, 67, 9, 12, 6, 45, 11, 7, 23, 40}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//n3
	// nums := [10]int{}
	// var n int
	// for i := 0; i < len(nums); i++ {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	fmt.Println(i, "-", nums[i])
	// }

	//n4
	// nums := [10]int{}
	// var n int
	// for i := 0; i < len(nums); i++ {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	if nums[i] >= 0 {
	// 		fmt.Println(nums[i])
	// 	}
	// }

	//n5
	// nums := [10]int{}
	// var n int
	// for i := 0; i < len(nums); i++ {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	if nums[i]%2 == 0 || nums[i] == 0 {
	// 		fmt.Println(nums[i])
	// 	}
	// }

	//n6
	// nums := [6]int{}
	// var n int
	// for i := 0; i < len(nums); i++ {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n * n
	// 	}
	// }
	// fmt.Println(nums)

	//n7
	// nums := [8]int{}
	// var n int
	// for i := 0; i < len(nums); i++ {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	//		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	if nums[i]%2 == 0 {
	// 		fmt.Println(nums[i])
	// 	}
	// }

	//n8
	// nums := [8]int{}
	// var n int
	// var count int
	// for i := 0; i < len(nums); i++ {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	if nums[i] >= 0 {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	//n9
	// nums := [8]int{}
	// var n int
	// var result int
	// for i := 0; i < len(nums); i++ {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	result += nums[i]
	// }
	// fmt.Println(result)

	//n10
	// nums := [6]int{}
	// var n int
	// result := 1
	// for i := range len(nums) {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	fmt.Println(nums)
	// 	if nums[i] != 0 {
	// 		result *= nums[i]
	// 	}
	// }
	// fmt.Println(result)

	//n11
	// nums := [10]int{}
	// var n, k int
	// //fmt.Scan(&k)
	// for i := range len(nums){
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	// if nums[i]%k == 0 {
	// 	// 	fmt.Println(nums[i])
	// 	// }
	// }
	// fmt.Scan(&k)
	// fmt.Println()
	// for i := range len(nums) {
	// 	if nums[i]%k == 0 {
	// 		fmt.Println(nums[i])
	// 	}
	// }

	//n12
	// nums := [10]int{}
	// var n, max int
	// for i := range len(nums){
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	if i == 0 {
	// 		max = n
	// 	} else if max < n {
	// 		max = n
	// 	}
	// }
	// fmt.Println(max)

	//n13
	// nums := [10]int{}
	// var n, min int
	// for i := range len(nums) {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	if i == 0 {
	// 		min = n
	// 	} else if min > n {
	// 		min = n
	// 	}
	// }
	// fmt.Println(min)

	//n14
	// nums := [10]int{}
	// var n, min, max int
	// for i := range len(nums) {
	// 	_, err := fmt.Scan(&n)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 		return
	// 	} else {
	// 		nums[i] = n
	// 	}
	// 	if i == 0 {
	// 		max = n
	// 		min = n
	// 	} else {
	// 		if max < n {
	// 			max = n
	// 		}
	// 		if min > n {
	// 			min = n
	// 		}
	// 	}
	// }
	// fmt.Println(min, " ", max)

	//n15
	nums := [10]int{}
	var n, m int
	var isSimilar bool
	for i := range len(nums) {
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Error")
			return
		} else {
			nums[i] = n
		}
	}
	fmt.Scan(&m)
	for i := range len(nums) {
		if nums[i] == m {
			isSimilar = true
			break
		} else {
			isSimilar = false
		}
	}
	if isSimilar {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	//lesstask1
	// var nums_size int
	// fmt.Scan(&nums_size)

	// nums := make([]int, nums_size)

	// for i := range nums_size {
	// 	var temp int
	// 	fmt.Scan(&temp)
	// 	nums[i] = temp
	// }

	// for i := range nums_size {
	// 	var temp_num int = nums[i]
	// 	for j := range nums_size {
	// 		if temp_num == nums[j] && i != j {
	// 			fmt.Println("Есть повторение!!!")
	// 			return
	// 		}
	// 	}
	// }
	// fmt.Println("Повторений нет!")

}
