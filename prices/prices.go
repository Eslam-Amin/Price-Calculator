package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct{
	IOManager iomanager.IOManager `json:"-"`
	TaxRate float64 `json:"taxRate"`
	InputPrices []float64 `json:"inputPrices"`
	TaxIncludedPrices map[string]string `json:"taxIncludedPrices"`
}



func (job *TaxIncludedPriceJob)Process(donechan chan bool) {
	job.LoadData()
	
	// if err != nil {
	// 	return err
	// }
	
	result := make(map[string]string)
		for _, price := range job.InputPrices{
			taxIncludedPrice := price * (1 + job.TaxRate)
			result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
		}

		job.TaxIncludedPrices = result
		job.IOManager.WriteResult(job)
		donechan <- true
}



func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}
	prices, err :=conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	job.InputPrices = prices
	return nil
}

func NewTaxIncludedPriceJob (ioManager iomanager.IOManager,taxRate float64) *TaxIncludedPriceJob{

	return &TaxIncludedPriceJob{
		IOManager: ioManager,
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}