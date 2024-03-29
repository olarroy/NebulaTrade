package mathnebula

import "math"

//Round -
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

//ToFixed -
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

//PercentOf -
func PercentOf(num float64, percent float64) float64 {

	return (percent * num) / 100
}
