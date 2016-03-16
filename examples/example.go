package main

import "github.com/sajari/regression"

func main() {
	var r regression.Regression
	r.SetObservedName("Z")
	r.SetVarName(0, "A")
	r.SetVarName(1, "B")
	r.SetVarName(2, "C")
	// Z = 12 + 1A + 2B +3C
	r.AddDataPoint(regression.DataPoint{Observed: 12, Variables: []float64{0, 0, 0}})
	r.AddDataPoint(regression.DataPoint{Observed: 13, Variables: []float64{1, 0, 0}})
	r.AddDataPoint(regression.DataPoint{Observed: 14, Variables: []float64{0, 1, 0}})
	r.AddDataPoint(regression.DataPoint{Observed: 15, Variables: []float64{0, 0, 1}})
	r.AddDataPoint(regression.DataPoint{Observed: 14, Variables: []float64{2, 0, 0}})
	r.AddDataPoint(regression.DataPoint{Observed: 16, Variables: []float64{0, 2, 0}})
	r.AddDataPoint(regression.DataPoint{Observed: 18, Variables: []float64{0, 0, 2}})

	r.RunLinearRegression()
	r.Dump(true)
}
