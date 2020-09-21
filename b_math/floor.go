package b_math

import "math"

func FloorF(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

func CeilF(x float32) float32 {
	return float32(math.Ceil(float64(x)))
}
