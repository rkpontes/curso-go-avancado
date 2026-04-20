package main

import "fmt"

func main() {
	fmt.Println("--- Copia Insegura ---")
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = slice1
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println()

	slice1[0] = 10
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println()

	slice2 = append(slice2, 6)
	slice2[0] = 99
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println()

	fmt.Println("--- Copia Segura ---")
	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, len(slice3))
	_ = copy(slice4, slice3)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println()

	slice3[0] = 10
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println()

}
