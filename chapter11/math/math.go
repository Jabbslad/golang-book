package math

// Finds the average value from a series of numbers
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Finds the minimum value from a series of numbers
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	min := xs[0]
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return min
}

// Finds the maximum value from a series of numbers
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	
	max := xs[0]
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}