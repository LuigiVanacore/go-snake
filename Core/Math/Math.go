package Math

func Min(x float32, y float32) float32 {
	if x == 0 || x == y {
		return x
	}
	if x < y {
		return x
	}
	return y
}

func Max(x float32, y float32) float32 {
	if x == 0 || x == y {
		return x
	}
	if x > y {
		return x
	}
	return y
}
