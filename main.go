package main

import (
	"fmt"

	"github.com/ezaz-ahmed/money-minder/filemanager"
	"github.com/ezaz-ahmed/money-minder/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTextIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
