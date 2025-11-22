package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct{
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64
}


func (job *TaxIncludedPriceJob)Process(){
	job.LoadData()
	result := make(map[string]float64)
		for _, price := range job.InputPrices{
			result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
		}

		fmt.Println(result)
}

func (job TaxIncludedPriceJob) LoadData(){
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Couldn't open file!")
		fmt.Println(err)
		return 
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil{
		fmt.Println("Couldn't Read file content!")
		fmt.Println(err)
		file.Close()
		return 
	}

	for lineIndex, line := range lines{
		 floatPrice , err := strconv.ParseFloat(line, 64)
		 if err != nil {
			fmt.Println("Converting Price to float failed!")
			fmt.Println(err)
			file.Close()
			return
		}
		job.InputPrices[lineIndex] = floatPrice
	}

}

func NewTaxIncludedPriceJob (taxRate float64) *TaxIncludedPriceJob{

	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}