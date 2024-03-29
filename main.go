package main

import "fmt"

func main() {
	num := 807
	arr := make([]int, 0)
	for num != 0 {
		arr = append(arr, num%10)
		num = num / 10
	}

	fmt.Println(arr)
}
