package main

import "math"

func normalizeMax(oldMin, oldMax, newMin, newMax, data float64) float64 {
	OldRange := (oldMax - oldMin)
	if OldRange == 0 {
		return newMin
	}
	NewRange := (newMax - newMin)
	return (((data - oldMin) * NewRange) / OldRange) + newMin
}

func powerLaw(max, c, velar, data float64) float64 {
	temp := data / max
	temp = math.Pow(temp, velar)
	return c * temp * max
}

func linearTransform(a, b, data float64) float64 {
	return a*data + b
}

func limitFuncTransform(low, high, limit, data float64) float64 {
	if data >= limit {
		return high
	} else {
		return low
	}
}
