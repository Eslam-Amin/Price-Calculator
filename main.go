package main

import "fmt"

func main() {
	var prices []float64 = []float64{10, 20, 30}
	taxRatess := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64 )
	for _, taxRate := range taxRatess{
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices{
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}

		result[taxRate] = taxIncludedPrices
	}

	fmt.Print(result)

}