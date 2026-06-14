package services

import (
	"errors"
	"math"
)

type QRResult struct {
	Q [][]float64
	R [][]float64
}

// FactorizeQR descompone una matriz en Q y R usando Gram-Schmidt
func FactorizeQR(matrix [][]float64) (QRResult, error) {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return QRResult{}, errors.New("matriz invalida")
	}

	cols := len(matrix[0])
	for i := range matrix {
		if len(matrix[i]) != cols {
			return QRResult{}, errors.New("la matriz debe ser rectangular")
		}
	}

	rows := len(matrix)
	k := rows
	if cols < k {
		k = cols
	}

	q := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		q[i] = make([]float64, k)
	}
	r := make([][]float64, k)
	for i := 0; i < k; i++ {
		r[i] = make([]float64, cols)
	}

	for j := 0; j < k; j++ {
		// copia la columna j de la matriz original
		v := make([]float64, rows)
		for i := 0; i < rows; i++ {
			v[i] = matrix[i][j]
		}

		// resta la proyección sobre los vectores anteriores
		for i := 0; i < j; i++ {
			var dot float64
			for t := 0; t < rows; t++ {
				dot += q[t][i] * matrix[t][j]
			}
			r[i][j] = dot
			for t := 0; t < rows; t++ {
				v[t] -= dot * q[t][i]
			}
		}

		// normaliza el vector para obtener la columna de Q
		var norm float64
		for _, val := range v {
			norm += val * val
		}
		norm = math.Sqrt(norm)
		if norm == 0 {
			return QRResult{}, errors.New("columnas linealmente dependientes")
		}

		r[j][j] = norm
		for t := 0; t < rows; t++ {
			q[t][j] = v[t] / norm
		}

		// calcula el resto de la fila j en R
		for col := j + 1; col < cols; col++ {
			var dot float64
			for t := 0; t < rows; t++ {
				dot += q[t][j] * matrix[t][col]
			}
			r[j][col] = dot
		}
	}

	return QRResult{Q: q, R: r}, nil

}
