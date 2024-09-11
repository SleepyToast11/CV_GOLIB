package grayscale

import "sort"

func zero(data []float64) float64 {
	return 0
}

func Medianer(data []float64) float64 {
	var median float64
	sort.Float64s(data)
	if len(data)%2 == 0 {
		median = (data[len(data)/2] + data[len(data)/2-1]) / 2
	} else {
		median = data[len(data)/2]
	}
	return median
}

func Averager(data []float64) float64 {
	if len(data) == 0 {
		panic(10)
	}
	var average float64
	for _, val := range data {
		average += val
	}
	return average / float64(len(data))
}

func WindowWithException(size int, exception func([]float64) float64, idxX, idxY int, data SingleImage) [][]float64 {
	windowSize := size*2 + 1
	vals := make([]float64, 0)
	for i := -size; i <= size; i++ {
		for j := -size; j <= size; j++ {
			val, err := data.GetVal(i+idxX, j+idxY)
			if err == nil {
				vals = append(vals, val)
			}
		}
	}
	exceptionVal := exception(vals)
	retVal := make([][]float64, windowSize)
	for i := range windowSize {
		retVal[i] = make([]float64, windowSize)
	}
	for i := -size; i <= size; i++ {
		for j := -size; j <= size; j++ {
			val, err := data.GetVal(i+idxX, j+idxY)
			if err == nil {
				retVal[j+size][i+size] = val
			} else {
				retVal[j+size][i+size] = exceptionVal
			}
		}
	}

	return retVal
}
