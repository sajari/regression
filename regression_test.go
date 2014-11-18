package regression

import (
	"testing"
)

func Test(t *testing.T) {
	var r Regression
	//r.Debug = true
	r.SetObservedName("Distance")
	r.SetVarName(0, "Weight")
	r.SetVarName(1, "Height")
	r.SetVarName(2, "Blood sugar")
	r.AddDataPoint(DataPoint{Observed: 98, Variables: []float64{483, 343, 0.0386}})
	r.AddDataPoint(DataPoint{Observed: 75, Variables: []float64{227, 419, 0.0705}})
	r.AddDataPoint(DataPoint{Observed: 0, Variables: []float64{380, 666, 0.0245}})
	r.AddDataPoint(DataPoint{Observed: 20, Variables: []float64{85, 833, 0.0567}})
	r.AddDataPoint(DataPoint{Observed: 50, Variables: []float64{188, 182, 0.0143}})
	r.AddDataPoint(DataPoint{Observed: 90, Variables: []float64{783, 1148, 0.0129}})
	r.RunLinearRegression()

	// TODO: This needs some real test data sets to test with.
	// Above data is not examined.
}
