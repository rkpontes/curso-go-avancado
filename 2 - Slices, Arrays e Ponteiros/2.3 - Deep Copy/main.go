package main

import "fmt"

func main() {
	fmt.Println("----- Copia Insegura --------")
	matriz1 := [][]int{{1, 2}, {3, 4}}
	matriz2 := make([][]int, len(matriz1))
	copy(matriz2, matriz1)
	matriz1[0][0] = 10
	fmt.Println(matriz1)
	fmt.Println(matriz2)
	fmt.Println()

	fmt.Println("----- Copia Segura --------")
	matriz3 := [][]int{{1, 2}, {3, 4}}
	matriz4 := copyMatriz(matriz3)
	matriz3[0][0] = 99
	fmt.Println(matriz3)
	fmt.Println(matriz4)
	fmt.Println()
}

func copyMatriz(matriz [][]int) [][]int {
	m2 := make([][]int, len(matriz))

	for i, slice := range matriz {
		m2[i] = make([]int, len(slice))
		copy(m2[i], slice)
	}

	return m2
}
