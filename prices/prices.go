package prices

import (
	"fmt"

	"github.com/ezaz-ahmed/money-minder/conversion"
	"github.com/ezaz-ahmed/money-minder/filemanager"
)

type TextIncludedPriceJob struct {
	TaxRate            float64
	InputPrices        []float64
	TextIncludedPrices map[string]string
}

func (job *TextIncludedPriceJob) LoadPrices() {
	const fileName = "prices.txt"

	lines, err := filemanager.ReadLines(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TextIncludedPriceJob) Process() {
	job.LoadPrices()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludePrice)
	}

	job.TextIncludedPrices = result
	filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTextIncludedPriceJob(taxRate float64) *TextIncludedPriceJob {
	return &TextIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
