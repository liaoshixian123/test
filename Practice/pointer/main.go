package main

import "fmt"

func main() {
	// var a int = 1
	// var b int = 2
	// p := a
	// d := b
	// fmt.Println("&a: ", &a, a)
	// fmt.Println("&b: ", &b, b)
	// fmt.Println("&p: ", &p, p)
	// fmt.Println("&d: ", &d, d)
	// fmt.Println("---------------------")

	// p = d
	// d = 5
	// fmt.Println("&a: ", &a, a)
	// fmt.Println("&b: ", &b, b)
	// fmt.Println("&p: ", &p, p)
	// fmt.Println("&d: ", &d, d)

	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr = append(arr, 5)
	arr = append(arr, 6)
	arrr(&arr)
}

func arrr(arr *[]int) {
	fmt.Println("arr地址: &arr", &arr)
	fmt.Println("arr地址轉變成值:*arr", *arr)
	fmt.Println("arr:", arr)

	brr := *arr
	crr := arr

	// &brr[] = 100000
	brr[3] = 1000
	fmt.Println("地址: ", &brr)
	// fmt.Println("地址轉變成值: ", *brr)
	fmt.Println("brr: ", brr)

	fmt.Println(&crr)

}
