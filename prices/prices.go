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

func (job *TextIncludedPriceJob) LoadPrices() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TextIncludedPriceJob) Process() error {
	err := job.LoadPrices()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludePrice)
	}

	job.TextIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

func NewTextIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TextIncludedPriceJob {
	return &TextIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
