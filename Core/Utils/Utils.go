package Utils

import "math"

func Min(x float64, y float64) float64 {
	if x == 0 || x == y {
		return x
	}
	if x < y {
		return x
	}
	return y
}

func Max(x float64, y float64) float64 {
	if x == 0 || x == y {
		return x
	}
	if x > y {
		return x
	}
	return y
}

func Distance(x, y, x2, y2 float64) float64 {
	dx := x - x2
	dy := y - y2
	ds := (dx * dx) + (dy * dy)
	return math.Sqrt(math.Abs(ds))
}
