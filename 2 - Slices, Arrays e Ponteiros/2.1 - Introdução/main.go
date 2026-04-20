package main

import "fmt"

func main() {
	var (
		array = [5]int{1, 2, 3, 4, 5}
		slice = []int{1, 2, 3, 4, 5}
	)

	fmt.Println("Array: ", array)
	fmt.Println("Tamanho do array: ", len(array))
	fmt.Println("Capacity do array: ", cap(array))

	fmt.Println("Slice: ", slice)
	fmt.Println("Tamanho do slice: ", len(slice))
	fmt.Println("Capacity do slice: ", cap(slice))

	fmt.Println()

	slice = append(slice, 6)

	fmt.Println("Slice: ", slice)
	fmt.Println("Tamanho do slice: ", len(slice))
	fmt.Println("Capacity do slice: ", cap(slice))

	printSlice(slice)

}

func printSlice(slice []int) {
	for _, i := range slice {
		fmt.Println(i)
	}
}
