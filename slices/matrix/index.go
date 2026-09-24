// slices can hold other slices, essentially creating a 2D matrix
// rows := [][]int{}
// Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers. The value of each cell is i * j where i and j are the indexes of the row and column respectively. Basically, we're building a multiplication chart.

package main

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		for j := range cols {
			matrix[i] = append(matrix[i], i*j)
		}
	}
	return matrix
}
