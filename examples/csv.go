/*
Loops through a CSV file and adds each row as a new data point in 
a linear regression calculation. Note: First row is assumed a header 
row with names of each variable. Secondly the first column is assumed 
to be your known observation. The remaining columns are variables.
*/
package main

import (
	"github.com/sajari/regression"
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
		)



func main() {
	// Initialise a new regression data set
	var r regression.Regression

    // Load the CSV data as data points
	dataFile, err := os.Open("chevy-mechanics.csv")
	if err != nil {
		fmt.Print(err, "\n")
	}
	reader := csv.NewReader(dataFile)
	row, err := reader.Read()
	count := 0
	for err == nil {
		if count == 0 {
			for i := 0; i < len(row); i++ {
				if i == 0 {
					r.SetObservedName(row[i])
				} else {
					r.SetVarName(i-1, row[i])
				}
			}
		} else {
			var d regression.DataPoint
			d.Variables = make([]float64, len(row)-1)
			for i := 0; i < len(row); i++ {
				if i == 0 {
					observed, _ := strconv.ParseFloat(row[i], 64);
					d.Observed = observed
				} else {
					variable, _ := strconv.ParseFloat(row[i], 64);
					d.Variables[i-1] = variable
				}
			}
			r.AddDataPoint(d)
		}
		count++
		row, err = reader.Read()
	}
	dataFile.Close()


	// Run the regression
    r.RunLinearRegression()
    r.Dump(false)
}