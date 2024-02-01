package prices

import (
	"fmt"

	"github.com/ezaz-ahmed/money-minder/conversion"
	"github.com/ezaz-ahmed/money-minder/iomanager"
)

type TextIncludedPriceJob struct {
	IOManager          iomanager.IOManager `json:"-"`
	TaxRate            float64             `json:"tax_rate"`
	InputPrices        []float64           `json:"input_prices"`
	TextIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TextIncludedPriceJob) LoadPrices() {
	lines, err := job.IOManager.ReadLines()

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
	job.IOManager.WriteResult(job)
}

func NewTextIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TextIncludedPriceJob {
	return &TextIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
