regression
=======

Multivariable Linear Regression in Go (golang)

installation
------------

    go get github.com/sajari/regression

usage
-----

Import the package, create a regression and add data to it. You can use as many variables as you like, in the below example there are 3 variables for each observation

    import "github.com/sajari/regression"

    func main() {
        var r regression.Regression
        r.SetObservedName("Distance")
        r.SetVarName(0, "Weight")
        r.SetVarName(1, "Height")
        r.SetVarName(2, "Blood sugar")
        r.AddDataPoint(regression.DataPoint{Observed : 98, Variables : []float64{483, 343, 0.0386}})
        r.AddDataPoint(regression.DataPoint{Observed : 75, Variables : []float64{227, 419, 0.0705}})
        r.AddDataPoint(regression.DataPoint{Observed : 0, Variables : []float64{380, 666, 0.0245}})
        r.AddDataPoint(regression.DataPoint{Observed : 20, Variables : []float64{85, 833, 0.0567}})
        r.AddDataPoint(regression.DataPoint{Observed : 50, Variables : []float64{188, 182, 0.0143}})
        r.RunLinearRegression()
        r.Dump()
    }

