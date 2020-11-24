package Math

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
