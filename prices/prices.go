package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TextIncludedPriceJob struct {
	TaxRate            float64
	InputPrices        []float64
	TextIncludedPrices map[string]float64
}

func (job *TextIncludedPriceJob) LoadPrices() {
	file, err := os.OpenFile("prices.txt", 22, 6088)

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading file content failed!")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for index, val := range lines {
		floatPrice, err := strconv.ParseFloat(val, 64)

		if err != nil {
			fmt.Println("Converting price to float failed!")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[index] = floatPrice
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

	fmt.Println((result))
}

func NewTextIncludedPriceJob(taxRate float64) *TextIncludedPriceJob {
	return &TextIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
