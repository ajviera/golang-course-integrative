package maths

import "errors"

// CalculateSuggestedPrice expose
func CalculateSuggestedPrice(slice []float64) (float64, float64, float64, error) {
	if len(slice) == 0 {
		return 0, 0, 0, errors.New("empty slice")
	}
	max := slice[0]
	min := slice[0]
	for _, val := range slice {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}
	suggested := Average(slice)
	return max, Round(suggested, 0.5), min, nil
}

// Average expose
func Average(slice []float64) float64 {
	total := 0.0
	for _, v := range slice {
		total += v
	}
	return total / float64(len(slice))
}

// Round expose
func Round(number, unit float64) float64 {
	return float64(int64(number/unit+0.5)) * unit
}
