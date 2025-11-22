package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct{
	IOManager filemanager.FileManager `json:"-"`
	TaxRate float64 `json:"taxRate"`
	InputPrices []float64 `json:"inputPrices"`
	TaxIncludedPrices map[string]string `json:"taxIncludedPrices"`
}



func (job *TaxIncludedPriceJob)Process(){
	job.LoadData()
	result := make(map[string]string)
		for _, price := range job.InputPrices{
			taxIncludedPrice := price * (1 + job.TaxRate)
			result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
		}

		job.TaxIncludedPrices = result
		job.IOManager.WriteResult(job)
}


func (job *TaxIncludedPriceJob) LoadData(){

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return 
	}
	prices, err :=conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return 
	}
	
	job.InputPrices = prices

}

func NewTaxIncludedPriceJob (fm filemanager.FileManager,taxRate float64) *TaxIncludedPriceJob{

	return &TaxIncludedPriceJob{
		IOManager: fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}