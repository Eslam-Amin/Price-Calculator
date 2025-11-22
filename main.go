package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)


func main() {
	taxRatess := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRatess))

	for index, taxRate := range taxRatess{
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		go priceJob.Process(doneChans[index])
		
		// if err != nil{
		// 	fmt.Println("Couldn't process the job")
		// 	fmt.Println(err)
		// }

		// cmd := cmdmanager.New()
		// priceJob = prices.NewTaxIncludedPriceJob(*cmd, taxRate)
		// priceJob.Process()
	}

	for _, doneChan := range doneChans{
		<-doneChan
	}
}