package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(tempArr(arr, 3))
	fmt.Println("abc")
}

//Array Rotation
func tempArr(arr []int, d int) []int {
	return append(arr[d:], arr[0:d]...)
}
