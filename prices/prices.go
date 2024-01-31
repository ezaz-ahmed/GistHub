package prices

import "fmt"

type TextIncludedPriceJob struct {
	TaxRate            float64
	InputPrices        []float64
	TextIncludedPrices map[string]float64
}

func (job TextIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println((result))
}

func NewTextIncludedPriceJob(taxRate float64) *TextIncludedPriceJob {
	return &TextIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
