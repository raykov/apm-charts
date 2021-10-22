package apmcharts

// Max returns the maximum float number from the slice
func Max(values []float64) (max float64) {
	for _, val := range values {
		if val > max {
			max = val
		}
	}

	return max
}

func getLegend(legend []string, index int) string {
	ln := len(legend)
	if ln == 0 || index >= ln {
		return ""
	}

	return legend[index]
}
