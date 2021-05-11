package math

// choice max value from arguments.
func Max(args ...interface{}) float64 {
	max := convertToFloat64(args[0])
	for _, a := range args[1:] {
		v := convertToFloat64(a)
		if v > max {
			max = v
		}
	}
	return max
}

// choice min value from arguments.
func Min(args ...interface{}) float64 {
	min := convertToFloat64(args[0])
	for _, a := range args[1:] {
		v := convertToFloat64(a)
		if v < min {
			min = v
		}
	}
	return min
}

// calculate summary.
func Sum(args ...interface{}) float64 {
	sum := 0.0
	for _, a := range args {
		sum += convertToFloat64(a)
	}
	return sum
}

// calculate avarage.
func Average(args ...interface{}) float64 {
	return Sum(args...) / float64(len(args))
}
