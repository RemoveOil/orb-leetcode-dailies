package binarymatrix

func countSquares(binaryMatrix [][]int) int {
	count := buildCountMatrix(binaryMatrix)
	I, J := len(count), len(count[0])
	for i := 1; i < I; i++ {
		for j := 1; j < J; j++ {
			minNeighbor := min(count[i-1][j-1], count[i][j-1], count[i-1][j])
			count[i][j] = minNeighbor*binaryMatrix[i][j] + binaryMatrix[i][j]
		}
	}

	var res int
	for i := 0; i < I; i++ {
		for j := 0; j < J; j++ {
			res += count[i][j]
		}
	}
	return res
}

func buildCountMatrix(binaryMatrix [][]int) [][]int {
	I, J := len(binaryMatrix), len(binaryMatrix[0])
	var ret [][]int
	for i := 0; i < I; i++ {
		ret = append(ret, make([]int, J))
	}

	for j := 0; j < J; j++ {
		ret[0][j] = binaryMatrix[0][j]
	}
	for i := 0; i < I; i++ {
		ret[i][0] = binaryMatrix[i][0]
	}
	return ret
}
