package cvgen

import "fmt"

func getWidth(maxWidth float64, ratios ...float64) ([]float64, error) {
	percentage := float64(0)
	res := make([]float64, 0, len(ratios))
	for _, ratio := range ratios {
		percentage += ratio
		res = append(res, (ratio/100)*maxWidth)
	}

	if int(percentage) != 100 {
		return res, fmt.Errorf("invalid ratio error, expect 100 but got %f", percentage)
	}

	return res, nil
}
