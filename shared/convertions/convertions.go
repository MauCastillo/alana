package convertions

import(
	"strconv"
)

// StringToFloat64 convert string to float64
func StringToFloat64(input string) (float64) {
	s, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return float64(0)
	}

	return s
}

// StringToInt64 convert string to int
func StringToInt64(input string) (int) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}

	return i
}
