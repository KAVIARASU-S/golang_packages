package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	arr := make([][]int, input)
	for i := 0; i < input; i++ {
		arr[i] = make([]int, i+1)
	}

	for i := 0; i < input; i++ {
		for j := 0; j <= i; j++ {
			if i == j || j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i-1][j-1]
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println(" ")
	}
}
