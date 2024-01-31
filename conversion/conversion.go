package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, val := range strings {
		floatPrice, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}
