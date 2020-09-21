package b_math

import "math"

func MinI(values ...int) int {
	m := math.MaxInt32
	for _, v := range values {
		if v < m {
			m = v
		}
	}
	return m
}

func MinF(values ...float32) float32 {
	m := float32(math.MaxFloat32)
	for _, v := range values {
		if v < m {
			m = v
		}
	}
	return m
}

func MaxI(values ...int) int {
	m := math.MinInt32
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}

func MaxF(values ...float32) float32 {
	m := float32(-math.MaxFloat32)
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}
