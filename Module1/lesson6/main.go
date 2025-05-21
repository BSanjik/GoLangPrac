package main

import "fmt"

func main() {
	//double Arrays
	//n1
	// doubleArrays := [][]int{
	// 	{12, 5, 7, 6},
	// 	{4, 0, 2, 7},
	// 	{9, 1, 3, 2},
	// 	{10, -2, 4, 6},
	// }
	// var n, m int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	fmt.Println(doubleArrays[n][m])
	// }

	//n2
	// var n, m, result int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range doubleArray {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 			result += doubleArray[i][j]
	// 		}
	// 	}
	// 	fmt.Println(result)
	// }

	//n3
	// var n, m, min int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 			if i == 0 && j == 0 {
	// 				min = doubleArray[i][j]
	// 			} else {
	// 				if min > doubleArray[i][j] {
	// 					min = doubleArray[i][j]
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(min)
	// }

	//n4
	// var n, m, max, indexI, indexJ int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 			if i == 0 && j == 0 {
	// 				max = doubleArray[i][j]
	// 			} else {
	// 				if max < doubleArray[i][j] {
	// 					max = doubleArray[i][j]
	// 					indexI, indexJ = i, j
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(max, indexI, indexJ)
	// }

	//n5
	// var n, m, result int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 			result += doubleArray[i][j]
	// 		}
	// 	}
	// 	fmt.Println(float64(result) / float64((n + m)))
	// }

	//n6
	// var n, m, x int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 		}
	// 	}
	// 	fmt.Scan(&x)
	// 	for i := range doubleArray {
	// 		for j := range doubleArray {
	// 			if doubleArray[i][j] == x {
	// 				fmt.Println("YES")
	// 				return
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("NO")
	// }

	//n7
	// var n, m, count int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 		}
	// 	}
	// }

	//n8
	// var n, m, count int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 			if doubleArray[i][j]%2 > 0 {
	// 				count++
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(count)
	// }

	//n9
	// var n, m int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 	}
	// 	for i := range n {
	// 		for j := range m {
	// 			if i == 0 {
	// 				fmt.Scan(&doubleArray[n-1][j])
	// 			} else {
	// 				if i == m-1 {
	// 					fmt.Scan(&doubleArray[0][j])
	// 				} else {
	// 					fmt.Scan(&doubleArray[i][j])
	// 				}
	// 			}
	// 		}
	// 	}
	// 	for i := range n {
	// 		for j := range m {
	// 			fmt.Print(doubleArray[i][j], " ")
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	//n10
	// var n, m int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 	}
	// 	for i := range n {
	// 		for j := range m {
	// 			if j == 0 {
	// 				fmt.Scan(&doubleArray[i][m-1])
	// 			} else {
	// 				if j == m-1 {
	// 					fmt.Scan(&doubleArray[i][0])
	// 				} else {
	// 					fmt.Scan(&doubleArray[i][j])
	// 				}
	// 			}
	// 		}
	// 	}
	// 	for i := range n {
	// 		for j := range m {
	// 			fmt.Print(doubleArray[i][j], " ")
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	//n11
	// var n, m int
	// var result bool = true
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 		}
	// 	}
	// 	for i := range n {
	// 		for j := range m {
	// 			if doubleArray[i][j] != doubleArray[i][j] {
	// 				result = false
	// 			}
	// 		}
	// 	}
	// 	if result {
	// 		fmt.Println("YES")
	// 	} else {
	// 		fmt.Println("NO")
	// 	}
	// }

	//n12
	// var n, m int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	result := make([]int, n)
	// 	doubleArray := make([][]int, n)
	// 	for i := range doubleArray {
	// 		doubleArray[i] = make([]int, m)
	// 		max := 0
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 			if i == 0 && j == 0 {
	// 				max = doubleArray[i][j]
	// 			} else {
	// 				if max < doubleArray[i][j] {
	// 					max = doubleArray[i][j]
	// 				}
	// 			}
	// 		}
	// 		result[i] = max
	// 	}
	// 	fmt.Println(result)
	// }

	//n13
	// var n, m, min int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	doubleArray := make([][]int, n)
	// 	result := make([]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 		}
	// 	}
	// 	for i := range n {
	// 		for j := range m {
	// 			if j == 0 {
	// 				min = doubleArray[i][j]
	// 			} else {
	// 				if min > doubleArray[i][j] {
	// 					min = doubleArray[i][j]
	// 				}
	// 			}
	// 		}
	// 		result[i] = min
	// 	}
	// 	fmt.Println(result)
	// }

	//n14
	// var n, m, sum int
	// _, err := fmt.Scan(&n, &m)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// } else {
	// 	result := make([]int, n)
	// 	doubleArray := make([][]int, n)
	// 	for i := range n {
	// 		doubleArray[i] = make([]int, m)
	// 		for j := range m {
	// 			fmt.Scan(&doubleArray[i][j])
	// 			sum += doubleArray[i][j]
	// 		}
	// 		result[i] = sum
	// 		sum = 0
	// 	}
	// 	fmt.Println(result)
	// }

	//n15
	var n, m, row, col, diagOne, diagSec, index int
	index = 0
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		fmt.Println("Error")
		return
	} else {
		result := make([]int, n+m+2)
		doubleArray := make([][]int, n)
		for i := range n {
			doubleArray[i] = make([]int, m)
			for j := range m {
				fmt.Scan(&doubleArray[i][j])
			}
		}
		for i := range n {
			for j := range m {
				row += doubleArray[i][j]
			}
			result[index] = row
			row = 0
			index++
		}
		for i := range n {
			for j := range m {
				col += doubleArray[i][j]
			}
			result[index] = col
			col = 0
			index++
		}
		for i := range n {
			diagOne += doubleArray[i][i]
		}
		result[index] = diagOne
		index++
		for i := len(doubleArray) - 1; i >= 0; i-- {
			diagSec += doubleArray[i][i]
		}
		result[index] = diagSec

		for i := range result {
			if result[0] != result[i] {
				fmt.Println("NO")
				return
			}
		}
		fmt.Println("YES")
	}

}
