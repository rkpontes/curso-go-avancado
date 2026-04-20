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
	matriz4 := make([][]int, len(matriz3))
	copyMatriz(matriz3, matriz4)
	matriz3[0][0] = 99
	fmt.Println(matriz3)
	fmt.Println(matriz4)
	fmt.Println()
}

func copyMatriz(m1 [][]int, m2 [][]int) {
	// percorre linha por linha
	for i := range m1 {
		m2[i] = make([]int, len(m1[i]))
		copy(m2[i], m1[i])
	}
}
