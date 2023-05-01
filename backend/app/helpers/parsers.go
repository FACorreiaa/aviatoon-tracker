package helpers

import (
	"fmt"
	"strconv"
)

func StringToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("error converting %s to int: %w", str, err)
	}
	return i, nil
}

func StringToFloat(str string) (float64, error) {
	if str == "" {
		return 0, nil
	}
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting %s to float: %w", str, err)
	}
	return i, nil
}
