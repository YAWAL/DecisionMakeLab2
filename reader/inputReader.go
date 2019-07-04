package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/YAWAL/DecisionMakeLab2/calculation"
)

const (
	decisionC  = "C"
	decisionB1 = "B1"
	decisionB  = "B"
	decisionA  = "A"
)

func readCSV(file string) (input [14]float64, err error) {
	f, err := os.Open(file)
	if err != nil {
		return [14]float64{}, err
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return [14]float64{}, err
	}
	for _, record := range records {
		for j, elem := range record {

			val, err := strconv.ParseFloat(elem, 64)
			if err != nil {
				return [14]float64{}, fmt.Errorf("can not parse float value from input file: %s", err)
			}
			input[j] = val
		}
	}
	return input, nil
}

func Input(file string) (dec1, dec2, dec4, dec5 calculation.Build, dec3 calculation.Wait, input [14]float64, err error) {
	input, err = readCSV(file)
	if err != nil {
		return calculation.Build{}, calculation.Build{}, calculation.Build{}, calculation.Build{}, calculation.Wait{}, input, err
	}
	fillData := func() {

		dec1.Name = decisionA
		dec1.Value = input[0]
		dec1.Profit = input[1]
		dec1.ProfitProbability = input[2]
		dec1.Loss = input[3]
		dec1.LossProbability = input[4]

		dec2.Name = decisionB
		dec2.Value = input[5]
		dec2.Profit = input[6]
		dec2.ProfitProbability = input[7]
		dec2.Loss = input[8]
		dec2.LossProbability = input[9]

		dec4.Name = "A1"
		dec4.Value = input[0]
		dec4.Profit = input[1]
		dec4.ProfitProbability = input[12]
		dec4.Loss = input[3]
		dec4.LossProbability = input[13]

		dec5.Name = decisionB1
		dec5.Value = input[5]
		dec5.Profit = input[6]
		dec5.ProfitProbability = input[12]
		dec5.Loss = input[8]
		dec5.LossProbability = input[13]

		dec3.Name = decisionC
		dec3.PositiveProbability = input[10]
		dec3.NegativeProbability = input[11]
		dec3.NewBuildA = dec4
		dec3.NewBuildB = dec5
	}
	fillData()
	return dec1, dec2, dec4, dec5, dec3, input, nil
}
