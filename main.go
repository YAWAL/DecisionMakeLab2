package main

import (
	"fmt"

	"github.com/YAWAL/DecisionMakeLab2/calculation"
	"github.com/YAWAL/DecisionMakeLab2/reader"
	"github.com/YAWAL/DecisionMakeLab2/reporter"
)

const (
	firstResult  = "1"
	secondResult = "2"
)

func main() {
	file := "input.csv"
	dec1, dec2, dec4, dec5, dec3, input, err := reader.Input(file)
	if err != nil {
		fmt.Errorf("error during preparing input data: %s", err)
	}

	results := make(map[string]string)
	results[dec1.Name] = fmt.Sprintf("%.2f", dec1.ProcessBuild())
	results[dec2.Name] = fmt.Sprintf("%.2f", dec2.ProcessBuild())
	results[dec3.Name] = fmt.Sprintf("%.2f", dec3.ProcessWait())
	results[dec4.Name] = fmt.Sprintf("%.2f", dec4.ProcessBuild())
	results[dec5.Name] = fmt.Sprintf("%.2f", dec5.ProcessBuild())
	results[firstResult] = calculation.Max(dec1.ProcessBuild(), dec2.ProcessBuild(), dec3.ProcessWait())
	results[secondResult] = calculation.Max(dec4.ProcessBuild(), dec5.ProcessBuild())

	fmt.Println(results)
	err = reporter.Report(results, input)
	if err != nil {
		fmt.Errorf("error during preparing input data: %s", err)
	}
}
