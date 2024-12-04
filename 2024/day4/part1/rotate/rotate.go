package rotate

// rotates a 2d array 45 degrees clockwise
func Rotate(matrix [][]rune, degrees int) [][]rune {
    

    var newMatrix [][]rune
    if degrees == 0 {
        return matrix
    }

    if degrees == 45 {

        diagonalCount := len(matrix) + len(matrix[0]) - 1
        newMatrix = make([][]rune, diagonalCount)
        for i := range newMatrix {
            newMatrix[i] = make([]rune, 0)
        }
        for i := 0; i < len(matrix); i++ {
            
            for j := 0; j < len(matrix); j++ {
                diagonalIndex := i + j
                newMatrix[diagonalIndex] = append(newMatrix[diagonalIndex], matrix[i][j])
            }
        }
        return newMatrix
    }
    
    if degrees == 90 {
        rowCount := len(matrix[0])
        colCount := len(matrix)
        newMatrix = make([][]rune, rowCount) // Number of rows in rotated matrix
        for i := range newMatrix {
            newMatrix[i] = make([]rune, colCount) // Number of columns in rotated matrix
        }

        // Populate the new matrix for a 90-degree rotation
        for i := 0; i < len(matrix); i++ {
            for j := 0; j < len(matrix[i]); j++ {
                newMatrix[j][colCount-1-i] = matrix[i][j]
            }
        }
        return newMatrix
    }
    return matrix
}
