package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)


func main() {
	taxRatess := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRatess{
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		err:=priceJob.Process()
		if err != nil{
			fmt.Println("Couldn't process the job")
			fmt.Println(err)
		}
		// cmd := cmdmanager.New()
		// priceJob = prices.NewTaxIncludedPriceJob(*cmd, taxRate)
		// priceJob.Process()
	}
}