package main

import "fmt"

func main() {
	fmt.Println("Closure Problem - 람다식 문제")
	slice := []func(){}
	for i := 0; i < 10; i++ {
		slice = append(slice, func() {
			fmt.Println(i)
		})
	}
	for _, v := range slice {
		v()
	}
	fmt.Println("i 변수를 참조로 가져오기 때문에 모든 값이 똑같음.")
	slice_new := []func(){}
	for i := 0; i < 10; i++ {
		temp := i
		slice_new = append(slice_new, func() {
			
			fmt.Println(temp)
		})
	}
	for _, v := range slice_new {
		v()
	}
	fmt.Println("i 변수를 temp 변수에 할당 시킨 후 temp 변수를 가져오기 때문에 값이 다름.")
}