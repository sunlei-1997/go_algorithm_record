package main

import "fmt"

func main() {
	res := make([]int,0)
	res = append(res, 1)
	fmt.Println(res)
	res = append([]int{2}, res...)
	fmt.Println(res)
}
