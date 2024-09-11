package main

import "math"

func filterApplier(mask, data [][]float64) float64 {
	if len(data) == len(mask) || len(mask[0]) == len(data[0]) {
		panic("mask not same size as window")
	}
	size := float64(len(mask[0]) * len(mask))
	var sum float64 = 0
	for y, yVal := range data {
		for x, val := range yVal {
			sum += mask[y][x] * val
		}
	}
	return sum / size
}

func sobelFilterHor(c float64, data [][]float64) float64 {
	if len(data) != 3 || len(data[0]) != 3 {
		panic("sobel filter require 3x3 window size")
	}
	mask := [][]float64{{-1, 0, 1}, {-c, 0, c}, {-1, 0, 1}}
	var sum float64 = 0
	for y, yVal := range data {
		for x, val := range yVal {
			sum += mask[y][x] * val
		}
	}
	return math.Abs(sum / 9)
}

func sobelFilterVer(c float64, data [][]float64) float64 {
	if len(data) != 3 || len(data[0]) != 3 {
		panic("sobel filter require 3x3 window size")
	}
	mask := [][]float64{{-1, c, 1}, {0, 0, 0}, {-1, c, 1}}
	var sum float64 = 0
	for y, yVal := range data {
		for x, val := range yVal {
			sum += mask[y][x] * val
		}
	}
	return math.Abs(sum / 9)
}

func sobelFilter(c float64, data [][]float64) float64 {
	return math.Sqrt(math.Pow(sobelFilterVer(c, data), 2) + math.Pow(sobelFilterHor(c, data), 2))
}

func pewitFilterHor(data [][]float64) float64 {
	return sobelFilterHor(1, data)
}

func pewitFilterVer(data [][]float64) float64 {
	return sobelFilterVer(1, data)
}

func pewitFilter(data [][]float64) float64 {
	return sobelFilter(1, data)
}

// like windows, mask is anchored at the middle and the size is the radius
// Precalculates the gaussian to save on computations
func laplacianOfGaussian(size int, stdDev float64) func(data [][]float64) float64 {
	windowSize := size*2 + 1
	preCalcMask := make([][]float64, windowSize)
	for i := range windowSize {
		preCalcMask[i] = make([]float64, windowSize)
		for j := range preCalcMask[i] {
			x := float64(j - size)
			y := float64(i - size)
			temp := -1 / (math.Pi * math.Pow(stdDev, 4))
			temp2 := 1 - ((math.Pow(x, 2) + math.Pow(y, 2)) / (2 * math.Pow(stdDev, 2)))
			temp3 := math.Exp(-((math.Pow(x, 2) + math.Pow(y, 2)) / (2 * math.Pow(stdDev, 2))))
			preCalcMask[i][j] = temp * temp2 * temp3
		}
	}
	return func(data [][]float64) float64 {
		return filterApplier(preCalcMask, data)
	}
}
