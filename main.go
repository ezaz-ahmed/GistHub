package main

import (
	"fmt"

	"github.com/ezaz-ahmed/money-minder/cmdmanager"
	"github.com/ezaz-ahmed/money-minder/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		cmdm := cmdmanager.New()

		priceJob := prices.NewTextIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}
}
