package main

import (
	"example.com/price-calculator/prices"
)


func main() {
	taxRatess := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRatess{
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}