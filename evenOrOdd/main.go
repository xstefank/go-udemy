package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range nums {
		fmt.Print(num)
		if num%2 == 0 {
			fmt.Println(" is even")
		} else {
			fmt.Println(" is odd")
		}
	}
}
